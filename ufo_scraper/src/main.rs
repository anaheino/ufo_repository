use scraper::Selector;

fn main() {

    let ufo_response = reqwest::blocking::get("https://nuforc.org/webreports/ndxpost.html")
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

}
