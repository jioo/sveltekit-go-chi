import { error, fail } from '@sveltejs/kit';
import * as api from '$lib/api';

export async function load({ locals, params }) {
    const albumID = +params.albumID;
    if (!albumID) return {}

    const album = await api.get(`albums/${albumID}`, locals.token);
    return album
}

export const actions = {
	save: async ({ request, locals, params }) => {
        let result;
		const form = await request.formData();
        const albumID = +params.albumID;
        const title = (form.get('title') !== 'undefined') ? form.get('title') : '';
        const artist = (form.get('artist') !== 'undefined') ? form.get('artist') : '';
        const price = (form.get('price') !== 'undefined') ? parseFloat(form.get('price')) : 0;
        const body = { title, artist, price }

        try {
            // save new album
            if (!albumID) {
                result = await api.post(`albums`, body, locals.token);
    
            // update existing album
            } else {
                result = await api.put(`albums/${albumID}`, body, locals.token);
            }
        } catch (err) {
            error(err);
        }
        
        if (result.errors) {
            return fail(401, result);
        }
        return { success: true };
	}
};