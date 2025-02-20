import jwt from 'jsonwebtoken';
import { JWT_KEY } from '$env/static/private';

export function handle({ event, resolve }) {
	event.locals.token = null;
	event.locals.user = null;

	try {
		const token = event.cookies.get('jwt');
		if (token) {
			const user = jwt.verify(token, JWT_KEY);
			event.locals.user = user;
			event.locals.token = token;
		}

	} catch (error) {
		console.error(error);
		event.cookies.set('jwt', '', { path: '/' });
	}

	return resolve(event);
}
