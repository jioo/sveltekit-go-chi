import { error, redirect } from '@sveltejs/kit';
import * as api from '$lib/api';

export async function load({ locals, cookies }) {
	if (!locals.user) redirect(302, '/');
    const albums = await api.get(`albums`, locals.token);
    return { albums }
}

export const actions = {
    delete: async ({ request, locals, params }) => {
		const data = await request.formData();
        const id = +data.get('id');

        const result = await api.del(`albums/${id}`, locals.token);
        return result;
    }
};