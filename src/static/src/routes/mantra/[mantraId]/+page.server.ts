import { makeApiUrl } from '$lib/api';
import type { Mantra } from '$lib/types';
import { fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';

export const load = (async ({ params }) => {
	return {
		mantra: await (await fetch(makeApiUrl(`/api/mantras/${params.mantraId}`))).json()
	} as {
		mantra: Mantra;
	};
}) satisfies PageServerLoad;

export const actions: Actions = {
	update: async ({ request, params }) => {
		const data = await request.formData();

		const res = await fetch(makeApiUrl(`/api/mantras/${params.mantraId}`), {
			method: 'POST',
			body: data
		});

		if (!res.ok) {
			return fail(400);
		}
		throw redirect(303, '/');
	},
	remove: async ({ request, params }) => {
		const data = await request.formData();

		const res = await fetch(makeApiUrl(`/api/mantras/${params.mantraId}`), {
			method: 'DELETE'
		});

		if (!res.ok) {
			return fail(400);
		}
		throw redirect(303, '/');
	}
};
