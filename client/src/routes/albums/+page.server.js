import { error, redirect } from '@sveltejs/kit';
import * as api from '$lib/api';

export async function load({ cookies }) {
    const albums = await api.get(`albums`);
    return { albums }
}

export const actions = {
    delete: async ({ request, locals, params }) => {
		const data = await request.formData();
        const id = +data.get('id');

        const result = await api.del(`albums/${id}`);
        return result;
    }
};