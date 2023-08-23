use scraper::Selector;
use table_extract::Table;

struct Sighting {
    date: String,
    city: String,
    state: String,
    country: String,
    shape: String,
    duration: String,
    description: String,
    report_date: String,
}
fn main() {

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
    ufo_links.for_each(|x| println!("{}", x));
    */

    let ufo_response = reqwest::blocking::get("https://nuforc.org/webreports/ndxp230710.html")
        .unwrap()
        .text()
        .unwrap();
    let report_page = scraper::Html::parse_document(&ufo_response).html();
    let table =   table_extract::Table::find_first(&report_page).unwrap();
    let mut sightings: Vec<Sighting> = Vec::new();
    for row in &table {
        let row_slice = row.as_slice();
        sightings.push(
        Sighting {
            date: row_slice.get(0).unwrap_or(&"Unavailable".to_string()).to_string(),
            city: row_slice.get(1).unwrap_or(&"Unavailable".to_string()).to_string(),
            state: row_slice.get(2).unwrap_or(&"Unavailable".to_string()).to_string(),
            country: row_slice.get(3).unwrap_or(&"Unavailable".to_string()).to_string(),
            shape: row_slice.get(4).unwrap_or(&"Unavailable".to_string()).to_string(),
            duration: row_slice.get(5).unwrap_or(&"Unavailable".to_string()).to_string(),
            description: row_slice.get(6).unwrap_or(&"Unavailable".to_string()).to_string(),
            report_date: row_slice.get(7).unwrap_or(&"Unavailable".to_string()).to_string(),
        });
    }
    for sight in &sightings {
        println!("time: {}", sight.date);
        println!("city: {}", sight.city);
        println!("state: {}", sight.state);
        println!("country: {}", sight.country);
        println!("shape: {}", sight.shape);
        println!("duration: {}", sight.duration);
        println!("description: {}", sight.description);
        println!("report_date: {}", sight.report_date);
    }
}
