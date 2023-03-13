<style>
	hr {
		color: var(--sx-gray-300);
	}
</style>

<Breadcrumbs links={breadcrumbLinks} />

<h2>Edit "{data.mantra.text}"</h2>
<form use:enhance method="POST" action="?/update" class="f-column gap-4">
	<TextInput name="text" value={data.mantra.text}>Mantra</TextInput>
	<div class="f-row justify-content-end">
		<button class="primary">Update</button>
	</div>
</form>

<hr style="width: 100%" />

<form use:enhance method="POST" action="?/remove">
	<button>Delete Mantra</button>
</form>

<script lang="ts">
	import { enhance } from '$app/forms';
	import { TextInput, Breadcrumbs } from 'sheodox-ui';
	import type { PageData } from './$types';

	export let data: PageData;

	const breadcrumbLinks = [
		{
			text: 'Mantras',
			href: '/'
		},
		{
			text: data.mantra.text
		}
	];
</script>
