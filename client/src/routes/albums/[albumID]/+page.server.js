import { error, redirect } from '@sveltejs/kit';
import * as api from '$lib/api';

export async function load({ locals, params }) {
    const albumID = +params.albumID;
    if (!albumID) return {}

    const album = await api.get(`albums/${albumID}`, locals.token);
    return album
}

export const actions = {
	save: async ({ request, locals, params }) => {
		const data = await request.formData();
        const albumID = +params.albumID;
        let result;

        try {
            // save new album
            if (!albumID) {
                result = await api.post(`albums`, {
                    title: data.get('title'),
                    artist: data.get('artist'),
                    price: parseFloat(data.get('price'))
                }, locals.token);
    
            // update existing album
            } else {
                result = await api.put(`albums/${albumID}`, {
                    title: data.get('title'),
                    artist: data.get('artist'),
                    price: parseFloat(data.get('price'))
                }, locals.token);
            }

        } catch (err) {
            console.log(err)
            return error(err.status, err.message);
        }
        
        return result;
	}
};