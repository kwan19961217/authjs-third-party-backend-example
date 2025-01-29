<script lang="ts">
	import type { PageData } from './$types';
	let { data }: { data: PageData } = $props();
	let session = data?.session;
</script>

<div>
	{#if session}
		<p>Session:</p>
		<pre>{JSON.stringify(session, null, 2)}</pre>
	{:else}
		<p>No session</p>
	{/if}
</div>
