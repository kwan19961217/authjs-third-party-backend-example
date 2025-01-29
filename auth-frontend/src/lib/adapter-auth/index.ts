import type { Adapter, AdapterUser, AdapterAccount, AdapterSession } from '@auth/core/adapters';
import type { AxiosInstance } from 'axios';

export default function MyAdapter(client: AxiosInstance): Adapter {
	return {
		async createUser(user: Omit<AdapterUser, 'id'>): Promise<AdapterUser> {
			const data = await client
				.post('/users', user)
				.then((res) => res.data)
				.catch((err) => {
					throw err;
				});
			return data;
		},
		async getUser(id: string): Promise<AdapterUser | null> {
			const data = await client
				.get(`/users/${id}`)
				.then((res) => res.data)
				.catch((err) => {
					if (err.response?.status === 404) return null;
					throw err;
				});
			return data;
		},
		async getUserByEmail(email: string): Promise<AdapterUser | null> {
			const data = await client
				.get(`/users/email/${email}`)
				.then((res) => res.data)
				.catch((err) => {
					if (err.response?.status === 404) return null;
					throw err;
				});
			return data;
		},
		async getUserByAccount(
			providerAccountId: Pick<AdapterAccount, 'provider' | 'providerAccountId'>
		): Promise<AdapterUser | null> {
			const data = await client
				.get(`/users/account/${providerAccountId.provider}/${providerAccountId.providerAccountId}`)
				.then((res) => res.data)
				.catch((err) => {
					if (err.response?.status === 404) return null;
					throw err;
				});
			return data;
		},

		async updateUser(user: AdapterUser): Promise<AdapterUser> {
			const data = await client
				.put(`/users/${user.id}`, user)
				.then((res) => res.data)
				.catch((err) => {
					if (err.response?.status === 404) return null;
					throw err;
				});
			return data;
		},
		async linkAccount(account: AdapterAccount): Promise<undefined | null | AdapterAccount> {
			const data = await client
				.post('/accounts', account)
				.then((res) => res.data)
				.catch((err) => {
					throw err;
				});
			return data;
		},
		async createSession(session: {
			sessionToken: string;
			userId: string;
			expires: Date;
		}): Promise<AdapterSession> {
			// TODO
			console.log('createSession', session);
			return new Promise((resolve) => {
				resolve(session);
			});
		},
		async getSessionAndUser(
			sessionToken: string
		): Promise<{ session: AdapterSession; user: AdapterUser } | null> {
			// TODO
			console.log('getSessionAndUser', sessionToken);
			return new Promise((resolve) => {
				resolve(null);
			});
		},
		async updateSession(
			session: Partial<AdapterSession> & Pick<AdapterSession, 'sessionToken'>
		): Promise<AdapterSession | null | undefined> {
			// TODO
			console.log('updateSession', session);
			return new Promise((resolve) => {
				resolve(null);
			});
		},
		async deleteSession(sessionToken: string): Promise<void> {
			// TODO
			console.log('deleteSession', sessionToken);
			return new Promise((resolve) => {
				resolve();
			});
		}
	};
}
