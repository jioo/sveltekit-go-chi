import { error, redirect } from '@sveltejs/kit';
import * as api from '$lib/api';

export async function load({ cookies }) {
    const albums = await api.get(`albums`);
    return { albums }
}