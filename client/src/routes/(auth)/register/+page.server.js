import { error, fail, redirect } from '@sveltejs/kit';
import jwt from 'jsonwebtoken';
import * as api from '$lib/api';
import { JWT_KEY } from '$env/static/private';

export const actions = {
	register: async ({ request, cookies }) => {
        let result;
		const form = await request.formData();
        const firstName = (form.get('firstName') !== 'undefined') ? form.get('firstName') : '';
        const lastName = (form.get('lastName') !== 'undefined') ? form.get('lastName') : '';
        const username = (form.get('username') !== 'undefined') ? form.get('username') : '';
        const password = (form.get('password') !== 'undefined') ? form.get('password') : '';
        const body = { firstName, lastName, username, password }

        try {
            result = await api.post(`register`, body);
            if (result.errors) {
                return fail(401, result);
            }
        } catch (err) {
            error(err);
        }

		return { success: true };
	}
};