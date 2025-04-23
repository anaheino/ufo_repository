import type { UfoSighting } from '@/types/types';
import {backendUrl} from "@/constants/contants";

export const insertSighting = async (sighting: UfoSighting)=> {
    try {
        const response = await fetch(`${backendUrl}/sighting`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(sighting),
        });
        return await response.json();
    } catch (error) {
        console.error('Error:', error);
    }
}