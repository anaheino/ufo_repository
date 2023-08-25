use scraper::Selector;
use table_extract::Table;
use mongodb::{Client, options::{ClientOptions, ResolverConfig}};
use std::{env, time::Duration};
use async_std::task;
use std::error::Error;
use bson::{doc, to_document};
use tokio;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
struct Sighting {
    date: String,
    city: String,
    state: String,
    country: String,
    shape: String,
    duration: String,
    description: String,
    report_date: String
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {

    let client_uri = env::var("MONGODB_UFO_URI").expect("You must set the MONGODB_URI environment var!");
    let client = Client::with_uri_str(client_uri).await?;
    let sightings_collection = client.database("i_want_to_believe").collection::<Sighting>("sightings");
    // sightings_collection.delete_many(doc! {"report_date": "7/10/23"}, None).await?;
    // task::sleep(Duration::from_millis(5000)).await;
    // Ok(())

    /*let ufo_response = reqwest::blocking::get("https://nuforc.org/webreports/ndxpost.html")
        .unwrap()
        .text()
        .unwrap();
    let ufo_document = scraper::Html::parse_document(&ufo_response);
    let selector = Selector::parse("a").unwrap();
    let ufo_pages = ufo_document.select(&selector)
        .filter_map(|n| n.value().attr("href"));
    let ufo_links = ufo_pages
        .skip(1)
        .map(|x| format!("{}{}", "https://nuforc.org/webreports/", x));
    ufo_links.for_each(|x| println!("{}", x));*/


    let ufo_response = reqwest::get("https://nuforc.org/webreports/ndxp230710.html")
        .await?
        .text()
        .await?;
    let report_page = scraper::Html::parse_document(&ufo_response).html();
    let table =   table_extract::Table::find_first(&report_page).unwrap();
    let mut sightings: Vec<Sighting> = Vec::new();

    for row in &table {
        let row_slice = row.as_slice();
        let a_element = row_slice.get(0).unwrap_or(&"<a>Unavailable</a>".to_string()).to_string();
        let string_array: Vec<_> = a_element.split(">").collect();
        let start_string = string_array.get(1).unwrap();
        let date_content: Vec<_> = start_string.split("<").collect();
        sightings.push(
          Sighting {
              date: date_content.get(0).unwrap().to_string(),
              city: row_slice.get(1).unwrap_or(&"Unavailable".to_string()).to_string(),
              state: row_slice.get(2).unwrap_or(&"Unavailable".to_string()).to_string(),
              country: row_slice.get(3).unwrap_or(&"Unavailable".to_string()).to_string(),
              shape: row_slice.get(4).unwrap_or(&"Unavailable".to_string()).to_string(),
              duration: row_slice.get(5).unwrap_or(&"Unavailable".to_string()).to_string(),
              description: row_slice.get(6).unwrap_or(&"Unavailable".to_string()).to_string(),
              report_date: row_slice.get(7).unwrap_or(&"Unavailable".to_string()).to_string()
          });
    }
    sightings_collection.insert_many(sightings, None).await?;
    Ok(())

    /*for sight in &sightings {
        println!("time: {}", sight.date);
        println!("city: {}", sight.city);
        println!("state: {}", sight.state);
        println!("country: {}", sight.country);
        println!("shape: {}", sight.shape);
        println!("duration: {}", sight.duration);
        println!("description: {}", sight.description);
        println!("report_date: {}", sight.report_date);
    }*/
}
