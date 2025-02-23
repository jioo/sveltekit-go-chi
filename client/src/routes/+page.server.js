import { redirect } from '@sveltejs/kit';

export async function load({ locals, cookies }) {
    redirect(302, '/login');
}