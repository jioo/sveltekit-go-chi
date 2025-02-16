import { error, redirect } from '@sveltejs/kit';
import * as api from '$lib/api';

export async function load({ params }) {
    const albumID = +params.albumID;
    const album = await api.get(`albums/${albumID}`);

    if (!album) return error(404, 'Album not found');
    return album
}

export const actions = {
	save: async ({ request, locals, params }) => {
		const data = await request.formData();
        const albumID = +params.albumID;

		const body = await api.put(`albums/${albumID}`, {
            title: data.get('title'),
            artist: data.get('artist'),
            price: parseFloat(data.get('price'))
		});

		return body;
	}
};