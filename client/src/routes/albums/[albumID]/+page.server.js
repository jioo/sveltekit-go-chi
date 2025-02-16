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

		const body = await api.put(`album/${albumID}`, {
            title: data.get('title'),
            artist: data.get('artist'),
            year: data.get('year')
		});

		if (body.errors) {
			return fail(401, body);
		}

		redirect(307, '/');
	}
};