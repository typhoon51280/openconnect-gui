<script lang="ts">
	import { Button, Form, FormGroup, Input, Label } from 'sveltestrap';
	export let loginUrl: string = '';
	export let email: string = '';
	export let password: string = '';
	export let cmd: string = '';
	function login(e: MouseEvent) {
		console.log("login event", e);
		console.log("loginUrl", loginUrl);
		globalThis.login(loginUrl, email, password)
		.then((result) => {
			console.log("login result:", result);
			if (result) {
				cmd = "sudo openconnect  --protocol=nc --verbose --dump-http-traffic -C 'DSID=" + result + "' " + loginUrl;
			}
		})
	}
</script>

<main>
	<Form>
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
	<Button color="primary" on:click={login}>Login</Button>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>