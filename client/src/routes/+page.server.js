import { error, fail, redirect } from '@sveltejs/kit';
import jwt from 'jsonwebtoken';
import * as api from '$lib/api';
import { JWT_KEY } from '$env/static/private';

export const actions = {
	login: async ({ request, cookies }) => {
        let result;
		const form = await request.formData();
        const username = form.get('username');
        const password = form.get('password');

        try {
            result = await api.post(`login`, { username, password });
            if (result.errors) {
                return fail(401, result);
            }

        } catch (err) {
            error(err);
        }

        verifyJWT(cookies, result)
		return { success: true };
	}
};

const verifyJWT = async (cookies, token) => {
    try {
        const user = jwt.verify(token, JWT_KEY);
        cookies.set('jwt', token, { path: '/' });
    } catch (err) {
        error(400, 'Invalid token');
    }
}