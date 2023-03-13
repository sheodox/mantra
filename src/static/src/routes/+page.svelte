<style>
	td {
		padding: var(--sx-spacing-3);
	}
</style>

<form method="POST" action="?/add" use:enhance class="mb-6">
	<Fieldset legend="New Mantra" contentClasses="f-column gap-2">
		<TextInput name="text" value="">Mantra</TextInput>
		<button class="primary">Add</button>
	</Fieldset>
</form>

{#if data.mantras.length}
	<table>
		<thead>
			<tr>
				<th>Mantra</th>
			</tr>
		</thead>
		<tbody>
			{#each data.mantras as mantra}
				<tr>
					<td>
						<a href="/mantra/{mantra.id}" class="inline-link">
							{mantra.text}
						</a>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
{:else}
	<p>You don't have any mantras yet.</p>
{/if}

<script lang="ts">
	import { TextInput, Fieldset } from 'sheodox-ui';
	import { enhance } from '$app/forms';
	import type { PageData } from './$types';

	export let data: PageData;
</script>
