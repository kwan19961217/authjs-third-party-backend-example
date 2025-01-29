import MyAdapter from '$lib/adapter-auth';
import authClient from '$lib/client';
import { SvelteKitAuth } from '@auth/sveltekit';
import GitHub from '@auth/sveltekit/providers/github';

export const { handle } = SvelteKitAuth({
	providers: [GitHub],
	adapter: MyAdapter(authClient()),
	session: {
		strategy: 'jwt'
	},
	callbacks: {
		jwt({ token, user }) {
			if (user) {
				// User is available during sign-in
				token.id = user.id;
			}
			return token;
		},
		session({ session, token }) {
			session.user.id = token.id as string;
			return session;
		}
	}
});
