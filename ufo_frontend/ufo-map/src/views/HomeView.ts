import type { UfoSighting } from '@/types/types';

export const insertSighting = async (sighting: UfoSighting)=> {
    try {
        const response = await fetch('http://localhost:8080/sighting', {
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