<script lang="ts">
	import {
		Container,
	} from 'sveltestrap';
	import { onMount } from 'svelte';
	import type { ConnectionItem } from './Connection.svelte';
	import Connections from './Connections.svelte';
	import Menu from './Menu.svelte';
	export let loginUrl: string = '';
	export let email: string = '';
	export let password: string = '';
	export let cmd: string = '';
	// let connections: Array<ConnectionItem> = []

	function login() {
		console.log("loginUrl", loginUrl);
		globalThis.Login(loginUrl, email, password, false)
		.then((result) => {
			console.log("login result:", result);
			if (result) {
				cmd = "sudo openconnect  --protocol=nc --verbose --dump-http-traffic -C 'DSID=" + result + "' " + loginUrl;
			}
		})
	}

</script>

<main>
	<Container id="connectionsPanel">
		<Menu></Menu>
		<Connections></Connections>
	<!-- <Form>
		<FormGroup>
			<Label for="loginUrl">Url</Label>
			<Input
			  type="url"
			  name="loginUrl"
			  id="loginUrl"
			  bind:value={loginUrl}
			  placeholder="loginUrl" />
		</FormGroup>
		<FormGroup>
			<Label for="email">Email</Label>
			<Input
			  type="email"
			  name="email"
			  id="email"
			  bind:value={email}
			  placeholder="email" />
		</FormGroup>
		<FormGroup>
			<Label for="password">Password</Label>
			<Input
			  type="password"
			  name="password"
			  id="password"
			  bind:value={password}
			  placeholder="password" />
		  </FormGroup>
		<FormGroup>
			<Label for="cookie">DSID</Label>
			<Input name="cookie" plaintext bind:value="{cmd}" />
		  </FormGroup>
	</Form>
	<Button color="primary" on:click={login}>Login</Button> -->
	</Container>
</main>

<style>
</style>