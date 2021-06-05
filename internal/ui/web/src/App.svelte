<script lang="ts">
	import {
		Container,
	} from 'sveltestrap';
	import ConnectionPanel from './ConnectionPanel.svelte';
	import Menu from './Menu.svelte';
</script>

<main>
	<Container class="mt-5 pt-3">
		<Menu></Menu>
		<ConnectionPanel></ConnectionPanel>
	</Container>
</main>

<style>
</style>