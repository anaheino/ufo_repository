use std::{env, time::Duration};
use std::collections::btree_map::Entry;
use std::collections::HashMap;
use async_std::task;
use std::error::Error;
use std::fmt::format;
use std::fs::File;
use std::io::{BufWriter, Write};
use std::path::Path;
use bson::{doc};
use tokio;
use serde::{Serialize, Deserialize};
use rand::Rng;
use mongodb::Client;

#[derive(Serialize, Deserialize, Clone)]
struct Sighting {
    date: String,
    city: String,
    state: String,
    country: String,
    shape: String,
    duration: String,
    description: String,
    report_date: String,
    has_images: bool,
    link: String
}
#[derive(Serialize, Deserialize)]
struct SightingWithLocation {
    original_sightings: Vec<SightingWithLatAndLon>,
    city: String,
    state: String,
    country: String,
}
#[derive(Serialize, Deserialize)]
struct SightingWithLatAndLon {
    date: String,
    city: String,
    state: String,
    country: String,
    shape: String,
    duration: String,
    description: String,
    report_date: String,
    has_images: bool,
    link: String,
    latitude: f64,
    longitude: f64,
}
#[derive(Serialize, Deserialize)]
struct Location {
    lat: f64,
    lng: f64,
}
#[derive(Serialize, Deserialize)]
struct ViewPort {
    northeast: Option<Location>,
    southwest: Option<Location>,
}
#[derive(Serialize, Deserialize)]
struct GeometryInfo {
    location: Option<Location>,
    bounds: Option<ViewPort>,
    location_type: Option<String>,
    viewport: Option<ViewPort>,
}
#[derive(Serialize, Deserialize)]
struct PlusCode {
    compound_code: Option<String>,
    global_code: Option<String>,
}
#[derive(Serialize, Deserialize)]
struct AddressComponent {
    long_name: Option<String>,
    short_name: Option<String>,
    types: Option<Vec<String>>,
}
#[derive(Serialize, Deserialize)]
struct GoogleGeocodingInfo {
    address_components: Vec<AddressComponent>,
    formatted_address: Option<String>,
    geometry: Option<GeometryInfo>,
    place_id: Option<String>,
    partial_match: Option<bool>,
    plus_code: Option<PlusCode>,
    types: Option<Vec<String>>,
}
#[derive(Serialize, Deserialize)]
struct GoogleGeocodingResponse {
    results: Vec<GoogleGeocodingInfo>,
    status: String,
    error_message: Option<String>,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let client_uri = env::var("MONGODB_UFO_URI").expect("You must set the MONGODB_URI environment var!");
    let google_api_key = env::var("GOOGLE_API_KEY").expect("You must set the GOOGLE_API_KEY environment var!");

    let client = Client::with_uri_str(client_uri).await?;
    let sightings_collection = client.database("i_want_to_believe")
        .collection::<SightingWithLatAndLon>("sightings_with_coords");
    let failed_sightings_collection = client.database("i_want_to_believe")
        .collection::<SightingWithLatAndLon>("sightings_with_failed_coord_query");

    let json_file_path = Path::new("./assets/grouped_sightings_that_errored_previously.json");
    let file = File::open(json_file_path).expect("file not found");
    let mut grouped_sightings: Vec<SightingWithLocation> = serde_json::from_reader(file).expect("error while reading");

    println!("{}", grouped_sightings.len());
    let mut skip_amount = 0;
    // grouped_sightings.drain(0..skip_amount);
    for mut single_sighting_group in grouped_sightings {
        let mut sightings_with_coords: Vec<SightingWithLatAndLon> = Vec::new();
        let mut city = single_sighting_group.city.to_string();
        let mut state = single_sighting_group.state.to_string();
        let mut country = single_sighting_group.country.to_string();
        if !city.is_empty() {
            city = format!("{}{}", city.to_string().trim().replace(" ", "%20"), "%20".to_string());
        }
        if !state.is_empty() {
            state = format!("{}{}", state.to_string().trim().replace(" ", "%20"), "%20".to_string());
        }
        if !country.is_empty() {
            country = country.to_string().trim().replace(" ", "%20");
        }
        let query_url = format!("https://maps.googleapis.com/maps/api/geocode/json?key={}&address={}{}{}",
                                google_api_key, city, state, country);
        let geo_cords_response = reqwest::get(query_url)
            .await?;
        let mut found_coords = false;
        let mut latitude = 0.000;
        let mut longitude = 0.000;
        match geo_cords_response.status() {
            reqwest::StatusCode::OK => {
                match geo_cords_response.json::<GoogleGeocodingResponse>().await {
                    Ok(parsed) => {
                        println!("{}", parsed.status);
                        if parsed.status.eq_ignore_ascii_case("OK") && !parsed.results.is_empty() {
                            let res = &parsed.results.get(0).unwrap().geometry;
                            if let Some(res) = &res {
                                let location = &res.location;
                                if let Some(location) = &location {
                                    latitude = location.lat;
                                    longitude = location.lng;
                                    found_coords = true;
                                }
                            }
                        } else {
                            println!("{}", single_sighting_group.country.to_string());
                            println!("{}", single_sighting_group.city.to_string());
                            println!("{}", single_sighting_group.state.to_string());
                        }
                        println!("{}", parsed.results.len());
                    }
                    Err(_) => {
                        println!("Current skip amount: {}", skip_amount);
                        panic!("Uh oh! Something unexpected happened!");
                    },
                };
            }
            other => {
                println!("Uh oh! Something unexpected happened, the query failed. Current skip amount {}, info on other: {:?}", skip_amount, other);
            }
        };
        for sighting in  single_sighting_group.original_sightings{
            sightings_with_coords.push(SightingWithLatAndLon {
                date: sighting.date.to_string(),
                city: sighting.city.to_string(),
                state: sighting.state.to_string(),
                country: sighting.country.to_string(),
                shape: sighting.shape.to_string(),
                duration: sighting.duration.to_string(),
                description: sighting.description.to_string(),
                report_date: sighting.report_date.to_string(),
                has_images: sighting.has_images,
                link: sighting.link.to_string(),
                latitude: latitude.clone(),
                longitude: longitude.clone()
            });
        }
        skip_amount += 1;
        if found_coords {
            sightings_collection.insert_many(sightings_with_coords, None).await?;
        } else {
            failed_sightings_collection.insert_many(sightings_with_coords, None).await?;
        }
        println!("Currently has ran {} iterations", skip_amount);
        task::sleep(Duration::from_millis(25)).await;
    }
    Ok(())
}
