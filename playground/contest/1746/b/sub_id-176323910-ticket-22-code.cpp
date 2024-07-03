#include <bits/stdc++.h>
using namespace std;
#define ll long long

#define fastio()                      \
	ios_base::sync_with_stdio(false); \
	cin.tie(NULL);                    \
	cout.tie(NULL)

bool comp(pair<ll, ll> &a, pair<ll, ll> &b)
{
	if (a.second < b.second)
		return 1;
	else if (a.second == b.second)
	{
		if (a.first < b.first)
			return 1;
		return 0;
	}
	return 0;
}

// choose mid items such that sum is minimum

ll check(vector<pair<ll, ll>> &tmp, ll mid, ll l)
{
	ll n = tmp.size();
	if (mid == 1)
	{
		for (int i = 0; i < n; i++)
		{
			if (tmp[i].first <= l)
				return 1;
		}
		return 0;
	}
	// dp.clear();
	// dp.resize(n+3, vector<vector<ll>>(mid+3, vector<ll>(2, 1e12)));
	vector<vector<ll>> prev(mid + 2, vector<ll>(2, 1e9));
	for (ll i = n - 1; i >= 0; i--)
	{
		vector<vector<ll>> curr(mid + 2, vector<ll>(2, 1e9));
		for (int j = mid; j >= 0; j--)
		{
			for (int k = 1; k >= 0; k--)
			{
				ll ans = 1e9;
				ans = min(ans, prev[j][k]);

				if (k == 0 && j >= 1)
					ans = min(ans, tmp[i].first - tmp[i].second + prev[j - 1][1]);
				else
				{
					if (j > 1)
						ans = min(ans, tmp[i].first + prev[j - 1][1]);
					else
						ans = min(ans, tmp[i].first + tmp[i].second);
				}

				curr[j][k] = ans;
			}
		}
		prev = curr;
	}
	// cout<<dp[0][mid][0]<<endl;
	if (prev[mid][0] <= l)
		return 1;
	return 0;
}

int main()
{
	fastio();
	ll t;
	cin >> t;
	while (t--)
	{
		ll n, l;
		cin >> n >> l;
		vector<pair<ll, ll>> vp(n);
		for (int i = 0; i < n; ++i)
		{
			cin >> vp[i].first >> vp[i].second;
		}
		vector<pair<ll, ll>> tmp = vp;
		sort(vp.begin(), vp.end());
		sort(tmp.begin(), tmp.end(), comp);

		// for(int i = 0 ; i < n ; i++){
		//     cout<<tmp[i].first<<' '<<tmp[i].second<<endl;
		// }
		// cout<<endl;

		ll low = 1, high = n, ans = 0;
		while (low <= high)
		{
			ll mid = (low + high) / 2;
			if (check(tmp, mid, l))
			{
				ans = mid;
				low = mid + 1;
			}
			else
				high = mid - 1;
		}
		// check(tmp, 3, l);
		cout << ans << endl;
	}
}