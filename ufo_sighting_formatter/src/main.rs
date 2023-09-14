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
    link: String,
    latitude: f64,
    longitude: f64,
}
#[derive(Serialize, Deserialize)]
struct SightingWithLocation {
    original_sightings: Vec<Sighting>,
    city: String,
    state: String,
    country: String,
}
#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let mut grouped_sightings: HashMap<String, Vec<Sighting>>= HashMap::new();
    let mut sightings_to_be_written: Vec<SightingWithLocation> = Vec::new();
    let json_file_path = Path::new("./assets/i_want_to_believe.sightings_with_failed_coord_query.json");
    let file = File::open(json_file_path).expect("file not found");
    let sightings:Vec<Sighting> = serde_json::from_reader(file).expect("error while reading");
    println!("{}", sightings.len());
    for mut sighting in sightings {
        if sighting.city.to_string().contains("(") {
            let name_array: Vec<_> = sighting.city.split("(").collect();
            sighting.city = name_array.get(0).unwrap_or(&"Empty city").parse().unwrap();
        }
        let key = format!("{}_{}_{}", sighting.country.to_lowercase().trim(), sighting.state.to_lowercase().trim(), sighting.city.to_lowercase().trim());
        let values = grouped_sightings.entry(key).or_insert_with(||vec![]);
        values.push(sighting);
    }
    println!("{}", grouped_sightings.len());

    let mut keys: Vec<String> = grouped_sightings.clone().into_keys().collect();
    keys.sort_by(|a, b| a.to_lowercase().cmp(&b.to_lowercase()));
    for key in keys {
        let values = grouped_sightings.get(&key).unwrap();
        let first_sighting = values.get(0).unwrap();
        sightings_to_be_written.push(SightingWithLocation {
            city: first_sighting.city.to_string().trim().to_string(),
            state: first_sighting.state.to_string().trim().to_string(),
            country: first_sighting.country.to_string().trim().to_string(),
            original_sightings: values.to_vec(),
        });
    }
    println!("{}", sightings_to_be_written.len());
    let file = File::create("./assets/grouped_sightings_that_errored_previously.json")?;
    let mut writer = BufWriter::new(file);
    serde_json::to_writer(&mut writer, &sightings_to_be_written)?;
    writer.flush();
    Ok(())
}
