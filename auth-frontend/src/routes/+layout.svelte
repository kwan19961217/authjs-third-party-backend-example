<script lang="ts">
	import { signIn, signOut } from '@auth/sveltekit/client';
	import '../app.css';
	import Header from '../stories/Header.svelte';
	import type { LayoutData } from './$types';
	let { children, data }: { children: any; data: LayoutData } = $props();
	let user = data?.session?.user;
</script>

<Header {user} onLogin={() => signIn('github')} onLogout={() => signOut()} />
{@render children()}
