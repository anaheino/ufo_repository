import type { UfoSighting } from "@/types/types";

export const groupSightings = (array: UfoSighting[], keyFunc: (x: UfoSighting) => string, asDictionary?: boolean): UfoSighting[][] | {[key: string]: UfoSighting[]} => {
    const groupedJsons = array.reduce((result, item) => {
        const keyValue: string = keyFunc(item);
        if (!result[keyValue]) {
            result[keyValue] = [];
        }
        result[keyValue].push(item);
        return result;
    }, {} as {[key: string]: UfoSighting[]});
    return !asDictionary ? Object.values(groupedJsons) : groupedJsons;
}