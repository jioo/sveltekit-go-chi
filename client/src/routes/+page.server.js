import { error, redirect } from '@sveltejs/kit';
import jwt from 'jsonwebtoken';
import * as api from '$lib/api';
import { JWT_KEY } from '$env/static/private';

export const actions = {
	login: async ({ request, cookies }) => {
		const form = await request.formData();
        const username = form.get('username');
        const password = form.get('password');

        if (!username || !password) {
            return error(400, 'Username and password are required');
        }

        let token = await api.post(`login`, { username, password });

        try {
            const user = jwt.verify(token, JWT_KEY);
            cookies.set('jwt', token, { path: '/' });

        } catch (err) {
            console.log(err);
            return error(400, 'Invalid token');
        }

		return { success: true };
	}
};