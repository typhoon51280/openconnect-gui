import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		loginUrl: 'https://vpnssl.virtual.posteitaliane.it/NFEC',
		email: 'ntt3dr7@posteitaliane.it',
		password: 'Pegasus3',
	}
});

export default app;