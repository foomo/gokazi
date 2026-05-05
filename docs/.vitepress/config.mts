import { defineConfig } from 'vitepress'

export default defineConfig({
	title: 'gokazi',
	description: 'Daemonless process manager',
	lang: 'en-US',
	lastUpdated: true,
	appearance: 'dark',
	ignoreDeadLinks: true,
	base: '/gokazi/',
	sitemap: {
		hostname: 'https://foomo.github.io/gokazi',
	},
	themeConfig: {
		logo: '/logo.png',
		outline: [2, 4],
		nav: [
			{ text: 'Guide', link: '/guide/introduction' },
			{ text: 'Reference', link: '/reference/configuration' },
			{ text: 'Recipes', link: '/recipes/local-dev-servers' },
		],
		sidebar: [
				{
					text: 'Guide',
					items: [
						{ text: 'Introduction', link: '/guide/introduction' },
						{ text: 'Installation', link: '/guide/installation' },
						{ text: 'Your first task', link: '/guide/your-first-task' },
						{ text: 'Managing tasks', link: '/guide/managing-tasks' },
						{ text: 'Troubleshooting', link: '/guide/troubleshooting' },
					],
				},
				{
					text: 'Reference',
					items: [
						{ text: 'Configuration', link: '/reference/configuration' },
						{
							text: 'CLI',
							link: '/reference/cli/',
							collapsed: false,
							items: [
								{ text: 'gokazi', link: '/reference/cli/gokazi' },
								{ text: 'list', link: '/reference/cli/gokazi_list' },
								{ text: 'stop', link: '/reference/cli/gokazi_stop' },
								{ text: 'config', link: '/reference/cli/gokazi_config' },
								{ text: 'version', link: '/reference/cli/gokazi_version' },
								{ text: 'completion', link: '/reference/cli/gokazi_completion' },
							],
						},
					],
				},
				{
					text: 'Recipes',
					items: [
						{ text: 'Local dev servers', link: '/recipes/local-dev-servers' },
						{ text: 'Make / Just integration', link: '/recipes/make-just-integration' },
						{ text: 'Multi-source config', link: '/recipes/multi-source-config' },
					],
				},
			{
					text: 'Contributing',
					collapsed: false,
					items: [
						{ text: 'Guideline', link: '/CONTRIBUTING' },
						{ text: 'Code of conduct', link: '/CODE_OF_CONDUCT' },
						{ text: 'Security guidelines', link: '/SECURITY' },
					],
				}
		],
		socialLinks: [
			{ icon: 'github', link: 'https://github.com/foomo/gokazi' },
		],
		editLink: {
			pattern: 'https://github.com/foomo/gokazi/edit/main/docs/:path',
		},
		search: {
			provider: 'local',
		},
		footer: {
			message: 'Made with ♥ <a href="https://www.foomo.org">foomo</a> by <a href="https://www.bestbytes.com">bestbytes</a>',
		},
	},
	markdown: {
		theme: {
			light: 'catppuccin-latte',
			dark: 'catppuccin-frappe',
		},
	},
	head: [
		['meta', { name: 'theme-color', content: '#ffffff' }],
		['link', { rel: 'icon', href: '/logo.png' }],
		['meta', { name: 'author', content: 'foomo by bestbytes' }],
		['meta', { property: 'og:title', content: 'foomo/gokazi' }],
		[
			'meta',
			{
				property: 'og:image',
				content: 'https://github.com/foomo/gokazi/blob/main/docs/public/banner.png?raw=true',
			},
		],
		[
			'meta',
			{
				property: 'og:description',
				content: 'Daemonless process manager.',
			},
		],
		['meta', { name: 'twitter:card', content: 'summary_large_image' }],
		[
			'meta',
			{
				name: 'twitter:image',
				content: 'https://github.com/foomo/gokazi/blob/main/docs/public/banner.png?raw=true',
			},
		],
		[
			'meta',
			{
				name: 'viewport',
				content: 'width=device-width, initial-scale=1.0, viewport-fit=cover',
			},
		],
	],
})
