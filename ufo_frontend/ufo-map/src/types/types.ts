export interface UfoSighting {
    city: string;
    country: string;
    date: string;
    description: string;
    duration: string;
    has_images: boolean;
    latitude: number;
    link: string;
    longitude: number;
    report_date: string;
    shape: string;
    state: string;
    _id: string;
}
export interface DateRangeString {
    startDate?: string;
    endDate?: string;
}
export interface SearchTerms {
    searchTerm: string;
    dates?: DateRangeString;
}
export interface CoordPair {
    latitude: number;
    longitude: number;
}