import { makeApiUrl } from '$lib/api';
import type { Mantra } from '$lib/types';
import { fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';

export const load = (async ({ fetch }) => {
	return {
		mantras: await (await fetch(makeApiUrl('/api/mantras'))).json()
	} as {
		mantras: Mantra[];
	};
}) satisfies PageServerLoad;

export const actions: Actions = {
	add: async ({ request }) => {
		const data = await request.formData();

		const res = await fetch(makeApiUrl('/api/mantras'), {
			method: 'POST',
			body: data
		});

		if (!res.ok) {
			return fail(400);
		}
		throw redirect(303, '/');
	},
	update: async ({ request }) => {
		const data = await request.formData();

		const res = await fetch(makeApiUrl(`/api/mantras/${data.get('id')}`), {
			method: 'POST',
			body: data
		});

		if (!res.ok) {
			return fail(400);
		}
		throw redirect(303, '/');
	},
	remove: async ({ request }) => {
		const data = await request.formData();

		const res = await fetch(makeApiUrl(`/api/mantras/${data.get('id')}`), {
			method: 'DELETE'
		});

		if (!res.ok) {
			return fail(400);
		}
		throw redirect(303, '/');
	}
};
