<script lang="ts">
	import {
		Container,
	} from 'sveltestrap';
	import Connections from './Connections.svelte';
	import Menu from './Menu.svelte';
</script>

<main>
	<Container class="mt-5 pt-3">
		<Menu></Menu>
		<Connections></Connections>
	</Container>
</main>

<style>
</style>