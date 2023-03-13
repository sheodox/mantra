<style global lang="scss">
	@import '../../node_modules/@fortawesome/fontawesome-free/css/all.min.css';

	main {
		margin: 0 auto;
		max-width: 90vw;
		width: 30rem;
	}
</style>

<SheodoxUIStyles />
<svelte:head>
	<title>Mantra</title>
</svelte:head>

<Header appName="Mantra" />

<main class="f-column mt-2">
	<slot />
</main>

<script lang="ts">
	import { SheodoxUIStyles, Header } from 'sheodox-ui';
</script>
