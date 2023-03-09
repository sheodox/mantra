{#if data.mantras.length}
	<table>
		<thead>
			<tr>
				<th>Mantra</th>
				<th>Delete</th>
			</tr>
		</thead>
		<tbody>
			{#each data.mantras as mantra}
				<tr>
					<td>
						<form use:enhance method="POST" action="?/update" class="f-row">
							<input type="hidden" value={mantra.id} name="id" />
							<TextInput name="text" value={mantra.text}>Mantra</TextInput>
							<button>Update</button>
						</form>
					</td>
					<td>
						<form use:enhance method="POST" action="?/remove">
							<input type="hidden" name="id" value={mantra.id} />
							<button>Delete</button>
						</form>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
{:else}
	<p>You don't have any mantras yet.</p>
{/if}

<form method="POST" action="?/add" use:enhance>
	<Fieldset legend="New Mantra">
		<TextInput name="text" value="">Mantra</TextInput>
		<button class="primary">Add</button>
	</Fieldset>
</form>

<script lang="ts">
	import { TextInput, Fieldset } from 'sheodox-ui';
	import { enhance } from '$app/forms';
	import type { PageData } from './$types';

	export let data: PageData;
</script>
