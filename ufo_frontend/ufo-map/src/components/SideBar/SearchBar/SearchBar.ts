import type { DateRangeString, SearchTerms } from "@/types/types";
import { toRaw } from "vue";
import {backendUrl} from "@/constants/contants";

export const querySightings = async (searchTerm: string, dates: DateRangeString) => {
    if (searchTerm.length > 0) {
        const searchTerms: SearchTerms = {
            searchTerm: searchTerm,
            dates: toRaw(dates),
        };
        const searchParams = formSearchString(searchTerms);
        const response = await fetch(`${backendUrl}/search?searchTerm=${searchParams}`);
        return await response.json();
    }
};
export const formSearchString = (searchTerms: SearchTerms) => {
    let searchString = `${searchTerms.searchTerm}`;
    if (searchTerms.dates) {
        if (searchTerms.dates.startDate) {
            searchString = searchString.concat(`&startDate=${searchTerms.dates.startDate}`)
        }
        if (searchTerms.dates.endDate) {
            searchString = searchString.concat(`&endDate=${searchTerms.dates.endDate}`)
        }
    }
    return searchString;
}