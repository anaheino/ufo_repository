use scraper::{Html, Selector};
use table_extract::Table;
use mongodb::Client;
use std::{env, time::Duration};
use async_std::task;
use std::error::Error;
use bson::{doc, to_document};
use chrono::{Datelike, NaiveDate, NaiveDateTime};
use chronoutil::shift_years;
use tokio;
use serde::{Serialize, Deserialize};
use rand::Rng;


#[derive(Serialize, Deserialize)]
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

struct AElement {
    link: String,
    text: String
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {

    let client_uri = env::var("MONGODB_UFO_URI").expect("You must set the MONGODB_URI environment var!");
    let client = Client::with_uri_str(client_uri).await?;
    let sightings_collection = client.database("i_want_to_believe").collection::<Sighting>("sightings");
    println!("Clearing database...");
    sightings_collection.delete_many(doc! {}, None).await?;

    let ufo_response = reqwest::get("https://nuforc.org/webreports/ndxpost.html")
        .await?
        .text()
        .await?;
    let ufo_document = scraper::Html::parse_document(&ufo_response);
    let selector = Selector::parse("a").unwrap();
    let ufo_pages = ufo_document.select(&selector)
        .filter_map(|n| n.value().attr("href"));
    let ufo_links: Vec<_> = ufo_pages
        .skip(1)
        .map(|x| format!("{}{}", "https://nuforc.org/webreports/", x))
        .collect();

    println!("Found {} different report pages...", ufo_links.len());

    let sleep_milliseconds = generate_sleep_milliseconds();
    println!("Sleeping for {} milliseconds...", sleep_milliseconds);
    task::sleep(Duration::from_millis(sleep_milliseconds)).await;
    for link in &ufo_links {
        let sightings = fetch_page_links(link).await?;
        println!("Inserting {} sightings to mongoDb...", sightings.len());
        sightings_collection.insert_many(sightings, None).await?;
        let sleep_milliseconds = generate_sleep_milliseconds();
        println!("Sleeping for {} milliseconds...", sleep_milliseconds);
        task::sleep(Duration::from_millis(sleep_milliseconds)).await;
    }
    Ok(())
}
async fn fetch_page_links(link: &String) -> Result<Vec<Sighting>, Box<dyn Error>> {
    println!("Fetching current ufo link: {}...", link);
    let ufo_response = reqwest::get(link)
        .await?
        .text()
        .await?;
    let report_page = scraper::Html::parse_document(&ufo_response).html();
    let table =   table_extract::Table::find_first(&report_page).unwrap();
    let mut sightings: Vec<Sighting> = Vec::new();

    for row in &table {
        let row_slice = row.as_slice();
        let date_link = row_slice.get(0).unwrap_or(&"<a href=\"no-link-here\">0/0/70 00:00</a>".to_string()).to_string();
        let a_element: AElement = parse_a_element_link_and_content(&date_link);
        let mut date_time;
        let date_time_parse = NaiveDateTime::parse_from_str(
            &a_element.text,
            "%m/%d/%Y %R"
        );
        if date_time_parse.is_ok() {
            date_time = date_time_parse.unwrap();
        } else {
            date_time = NaiveDateTime::parse_from_str(
                "01/01/70 00:00",
                "%m/%d/%Y %R"
            ).unwrap();
        }
        date_time = correct_date(date_time);
        let report_date = row_slice.get(7).unwrap_or(&"0/0/70".to_string()).to_string();
        let mut report_date_time;
        let report_date_time_parse = NaiveDate::parse_from_str(
            &report_date.to_string(),
            "%m/%d/%Y"
        );
        if report_date_time_parse.is_ok() {
            report_date_time = report_date_time_parse.unwrap().and_hms_opt(0, 0, 0).unwrap();
        } else {
            report_date_time = NaiveDateTime::parse_from_str(
                "01/01/70 00:00",
                "%m/%d/%Y %R"
            ).unwrap();
        }
        report_date_time = correct_date(report_date_time);
        let full_link = format!("{}{}", "https://nuforc.org/webreports/", a_element.link);
        let has_images_content = row_slice.get(8).unwrap_or(&"".to_string()).to_string().to_lowercase();
        sightings.push(
            Sighting {
                date: date_time.format("%FT%T").to_string(),
                city: row_slice.get(1).unwrap_or(&"Unavailable".to_string()).to_string(),
                state: row_slice.get(2).unwrap_or(&"Unavailable".to_string()).to_string(),
                country: row_slice.get(3).unwrap_or(&"Unavailable".to_string()).to_string(),
                shape: row_slice.get(4).unwrap_or(&"Unavailable".to_string()).to_string(),
                duration: row_slice.get(5).unwrap_or(&"Unavailable".to_string()).to_string(),
                description: row_slice.get(6).unwrap_or(&"Unavailable".to_string()).to_string(),
                report_date: report_date_time.format("%FT%T").to_string(),
                has_images: has_images_content.eq("yes"),
                link: full_link
            });
    }

    Ok(sightings)
}
fn correct_date(mut date_time: NaiveDateTime) -> NaiveDateTime {
    if date_time.year() > 23 {
        date_time = shift_years(date_time, 1900);
    } else {
        date_time = shift_years(date_time, 2000);
    }
    return date_time;
}

fn generate_sleep_milliseconds() -> u64 {
    let mut rng = rand::thread_rng();
    return rng.gen_range(1500..3000);
}

fn parse_a_element_link_and_content(html_text: &str) -> AElement {
    let html= Html::parse_fragment(html_text);
    let sel = Selector::parse("a").unwrap();
    let link_list: Vec<_> = html.select(&sel).filter_map(|n| n.value().attr("href")).collect();
    return AElement {
        text: html.select(&sel).next().unwrap().inner_html(),
        link: link_list.get(0).unwrap().to_string()
    };
}