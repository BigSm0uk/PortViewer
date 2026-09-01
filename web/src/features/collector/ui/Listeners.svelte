<script lang="ts">
	import type { Listener } from '../contract';
	import * as Table from "$lib/components/ui/table/index.js";

	let { listeners = $bindable() }: { listeners: Listener[] } = $props();
</script>

<Table.Root>
 <Table.Caption>A list of your recent invoices.</Table.Caption>
 <Table.Header>
  <Table.Row>
   <Table.Head class="w-25">Process</Table.Head>
   <Table.Head>PID</Table.Head>
   <Table.Head>Protocol</Table.Head>
   <Table.Head>address</Table.Head>
   <Table.Head>port</Table.Head>
   <Table.Head>exe</Table.Head>
  </Table.Row>
 </Table.Header>
 <Table.Body>
  {#each listeners as listener, index (index)}
   <Table.Row>
    <Table.Cell class="font-medium">{listener.process}</Table.Cell>
    <Table.Cell>{listener.pid}</Table.Cell>
    <Table.Cell>{listener.protocol}</Table.Cell>
    <Table.Cell>{listener.address}</Table.Cell>
    <Table.Cell>{listener.port}</Table.Cell>
    <Table.Cell>{listener.exe}</Table.Cell>
   </Table.Row>
  {/each}
 </Table.Body>
</Table.Root>
