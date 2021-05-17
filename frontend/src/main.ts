import App from './App.svelte';

// bootstrap
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap-icons/font/bootstrap-icons.css';

const app = new App({
	target: document.body,
	props: {
		name: 'world'
	}
});

export default app;