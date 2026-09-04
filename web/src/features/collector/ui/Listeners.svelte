<script lang="ts">
	import ChevronDownIcon from '@lucide/svelte/icons/chevron-down';
	import SearchIcon from '@lucide/svelte/icons/search';
	import XIcon from '@lucide/svelte/icons/x';
	import { Badge } from '$lib/components/ui/badge/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import { Input } from '$lib/components/ui/input/index.js';

	import type { Listener } from '../contract';
	import {
		filterListeners,
		formatPortsCount,
		getListenerKey,
		groupListeners,
		summarizeListeners
	} from './listeners';

	let { listeners }: { listeners: Listener[] } = $props();
	let query = $state('');

	let filteredListeners = $derived(filterListeners(listeners, query));
	let groups = $derived(groupListeners(filteredListeners));
	let summary = $derived(summarizeListeners(groups));
	let totalServices = $derived(groupListeners(listeners).length);
</script>

<svelte:head>
	<title>PortViewer — открытые порты</title>
	<meta
		name="description"
		content="Сервисы и процессы, прослушивающие сетевые порты на этом устройстве."
	/>
</svelte:head>

<section class="min-h-full bg-[#f4f7fb] text-slate-950" aria-labelledby="listeners-title">
	<header class="border-b border-slate-200/90 bg-white px-4 py-5 sm:px-6 lg:px-8 lg:py-7">
		<div class="mx-auto flex max-w-7xl flex-col gap-5 lg:flex-row lg:items-end lg:justify-between">
			<div class="max-w-2xl">
				<div class="mb-3 flex items-center gap-2 text-sm font-medium text-slate-500">
					<span class="relative flex size-2.5" aria-hidden="true">
						<span
							class="absolute inline-flex size-full animate-ping rounded-full bg-sky-400 opacity-50"
						></span>
						<span class="relative inline-flex size-2.5 rounded-full bg-sky-600"></span>
					</span>
					Снимок сетевых слушателей
				</div>
				<h1 id="listeners-title" class="text-2xl font-semibold tracking-[-0.035em] sm:text-3xl">
					Сервисы и открытые порты
				</h1>
				<p class="mt-2 max-w-xl text-sm leading-6 text-slate-600 sm:text-base">
					Процессы, принимающие входящие соединения на этом устройстве.
				</p>
			</div>

			<label class="relative block w-full lg:max-w-md">
				<span class="sr-only">Найти сервис, порт, адрес или исполняемый файл</span>
				<SearchIcon
					class="pointer-events-none absolute top-1/2 left-3.5 size-4 -translate-y-1/2 text-slate-400"
					aria-hidden="true"
				/>
				<Input
					bind:value={query}
					type="search"
					placeholder="Сервис, PID, порт, адрес или путь"
					class="search-input h-11 rounded-lg border-slate-300 bg-white pr-11 pl-10 shadow-sm focus-visible:border-sky-600 focus-visible:ring-sky-100"
				/>
				{#if query}
					<Button
						type="button"
						variant="ghost"
						size="icon-sm"
						class="absolute top-1/2 right-2 -translate-y-1/2 text-slate-500 hover:bg-slate-100 hover:text-slate-950 focus-visible:ring-sky-200"
						onclick={() => (query = '')}
						aria-label="Очистить поиск"
					>
						<XIcon class="size-4" aria-hidden="true" />
					</Button>
				{/if}
			</label>
		</div>
	</header>

	<main class="mx-auto max-w-7xl px-4 py-5 sm:px-6 lg:px-8 lg:py-7">
		<Card.Root class="gap-0 rounded-xl border border-slate-200 py-0 shadow-xs ring-0">
			<Card.Content class="px-0">
				<dl class="grid grid-cols-2 overflow-hidden lg:grid-cols-4">
					<div class="border-r border-b border-slate-200 px-4 py-4 sm:px-5 lg:border-b-0">
						<dt class="text-sm text-slate-500">Сервисов</dt>
						<dd class="mt-1 text-2xl font-semibold tracking-tight tabular-nums">
							{summary.services}
						</dd>
					</div>
					<div class="border-b border-slate-200 px-4 py-4 sm:px-5 lg:border-r lg:border-b-0">
						<dt class="text-sm text-slate-500">Портов</dt>
						<dd class="mt-1 text-2xl font-semibold tracking-tight tabular-nums">{summary.ports}</dd>
					</div>
					<div class="border-r border-slate-200 px-4 py-4 sm:px-5">
						<dt class="flex items-center gap-2 text-sm text-slate-500">
							<span class="size-2 rounded-full bg-sky-500" aria-hidden="true"></span>
							TCP
						</dt>
						<dd class="mt-1 text-2xl font-semibold tracking-tight tabular-nums">{summary.tcp}</dd>
					</div>
					<div class="px-4 py-4 sm:px-5">
						<dt class="flex items-center gap-2 text-sm text-slate-500">
							<span class="size-2 rounded-full bg-violet-500" aria-hidden="true"></span>
							UDP
						</dt>
						<dd class="mt-1 text-2xl font-semibold tracking-tight tabular-nums">{summary.udp}</dd>
					</div>
				</dl>
			</Card.Content>
		</Card.Root>

		<div class="mt-7 flex items-baseline justify-between gap-4">
			<h2 class="text-base font-semibold">Сервисы</h2>
			<p class="text-sm text-slate-500" aria-live="polite">
				{#if query.trim()}
					Показано {summary.services} из {totalServices}
				{:else}
					{summary.services} всего
				{/if}
			</p>
		</div>

		{#if groups.length > 0}
			<div class="mt-3 space-y-3">
				{#each groups as group (group.key)}
					<Card.Root
						class="service-group gap-0 rounded-xl border border-slate-200 py-0 shadow-xs ring-0"
					>
						<Card.Header class="border-b border-slate-200 px-4 py-4 sm:px-5">
							<div class="flex min-w-0 items-start justify-between gap-4">
								<div class="flex min-w-0 items-center gap-3">
									<span
										class="grid size-10 shrink-0 place-items-center rounded-lg bg-slate-950 text-sm font-semibold text-white shadow-inner"
										aria-hidden="true"
									>
										{group.process.slice(0, 1).toLocaleUpperCase() || '?'}
									</span>
									<div class="min-w-0">
										<h3
											class="truncate text-base font-semibold tracking-[-0.015em]"
											title={group.process}
										>
											{group.process || 'Неизвестный процесс'}
										</h3>
										<p class="mt-0.5 text-xs text-slate-500">PID {group.pid}</p>
									</div>
								</div>
								<Badge variant="secondary" class="h-6 shrink-0 px-2.5 text-slate-600">
									{formatPortsCount(group.listeners.length)}
								</Badge>
							</div>

							{#if group.exe}
								<details class="exe-details mt-3 min-w-0">
									<summary
										class="flex min-w-0 cursor-pointer list-none items-center gap-1.5 text-xs text-slate-500 marker:hidden hover:text-slate-800 focus-visible:rounded-sm focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-sky-600"
										title={group.exe}
									>
										<span class="shrink-0">Исполняемый файл:</span>
										<code class="min-w-0 truncate font-mono text-[0.6875rem] text-slate-600">
											{group.exe}
										</code>
										<ChevronDownIcon
											class="exe-chevron size-3.5 shrink-0 transition-transform"
											aria-hidden="true"
										/>
									</summary>
									<code
										class="mt-2 block rounded-md bg-slate-50 px-3 py-2 font-mono text-[0.6875rem] leading-5 break-all text-slate-700 ring-1 ring-slate-200"
									>
										{group.exe}
									</code>
								</details>
							{:else}
								<p class="mt-3 text-xs text-slate-400">Путь к исполняемому файлу недоступен</p>
							{/if}
						</Card.Header>

						<Card.Content class="px-0">
							<ul class="divide-y divide-slate-100" aria-label={`Порты сервиса ${group.process}`}>
								{#each group.listeners as listener, index (getListenerKey(listener, index))}
									<li class="port-row grid gap-3 px-4 py-3.5 sm:px-5">
										<div>
											<p class="port-label text-[0.6875rem] font-medium text-slate-400">Порт</p>
											<p class="mt-0.5 font-semibold tabular-nums">{listener.port}</p>
										</div>
										<div>
											<p class="port-label text-[0.6875rem] font-medium text-slate-400">Протокол</p>
											<Badge
												variant="secondary"
												class={listener.protocol.toLocaleLowerCase() === 'tcp'
													? 'mt-1 bg-sky-100 text-sky-700'
													: listener.protocol.toLocaleLowerCase() === 'udp'
														? 'mt-1 bg-violet-100 text-violet-700'
														: 'mt-1'}
											>
												{listener.protocol.toLocaleUpperCase()}
											</Badge>
										</div>
										<div class="min-w-0">
											<p class="port-label text-[0.6875rem] font-medium text-slate-400">Адрес</p>
											<code
												class="mt-0.5 block truncate font-mono text-xs leading-6 text-slate-700"
												title={listener.address}
											>
												{listener.address || '—'}
											</code>
										</div>
									</li>
								{/each}
							</ul>
						</Card.Content>
					</Card.Root>
				{/each}
			</div>
		{:else}
			<div
				class="mt-3 rounded-xl border border-dashed border-slate-300 bg-white px-6 py-14 text-center"
			>
				<div
					class="mx-auto mb-4 grid size-11 place-items-center rounded-full bg-slate-100 text-slate-500"
				>
					<SearchIcon class="size-5" aria-hidden="true" />
				</div>
				<h3 class="font-semibold">{query.trim() ? 'Ничего не найдено' : 'Открытых портов нет'}</h3>
				<p class="mx-auto mt-2 max-w-sm text-sm leading-6 text-slate-500">
					{query.trim()
						? 'Проверьте запрос или попробуйте найти по другому полю.'
						: 'Когда сервис начнёт прослушивать порт, он появится в этом списке.'}
				</p>
				{#if query.trim()}
					<Button
						type="button"
						class="mt-5 bg-slate-950 text-white hover:bg-slate-800 focus-visible:ring-sky-200"
						onclick={() => (query = '')}
					>
						Очистить поиск
					</Button>
				{/if}
			</div>
		{/if}
	</main>
</section>

<style>
	:global(.service-group) {
		container-type: inline-size;
	}

	.port-row {
		grid-template-columns: minmax(5rem, 0.7fr) minmax(5.5rem, 0.8fr) minmax(0, 2fr);
		align-items: center;
	}

	:global(.search-input::-webkit-search-cancel-button) {
		appearance: none;
	}

	:global(.exe-details[open] .exe-chevron) {
		transform: rotate(180deg);
	}

	@container (max-width: 31rem) {
		.port-row {
			grid-template-columns: 1fr 1fr;
			align-items: start;
		}

		.port-row > :last-child {
			grid-column: 1 / -1;
		}
	}

	@media (prefers-reduced-motion: reduce) {
		:global(.animate-ping) {
			animation: none;
		}

		:global(.exe-chevron) {
			transition: none;
		}
	}
</style>
