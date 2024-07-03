// #include<bits/stdc++.h>
// using namespace std;
// #define ll long long

// #define fastio() ios_base::sync_with_stdio(false);cin.tie(NULL);cout.tie(NULL)

// bool comp(pair<ll,ll> &a, pair<ll,ll> &b){
//     if(a.second < b.second) return 1;
//     else if(a.second == b.second){
//         if(a.first < b.first) return 1;
//         return 0;
//     }
//     return 0;
// }

// // choose mid items such that sum is minimum
// vector<vector<ll>> pre;
// vector<vector<ll>> curr;

// ll check(vector<pair<ll,ll>> &tmp, ll mid, ll l){
//     ll n = tmp.size();
//     if(mid == 1){
//         for(int i = 0 ; i < n ; i++){
//             if(tmp[i].first <= l) return 1;
//         }
//         return 0;
//     }
//     // dp.clear();
//     // dp.resize(n+3, vector<vector<ll>>(mid+3, vector<ll>(2, 1e12)));
//     pre.clear();
//     pre.resize(mid+1, vector<ll> (2, 1e12));
//     for(ll i = n-1 ; i >= 0 ; i--){
//         curr.clear();
//         curr.resize(mid+1, vector<ll> (2, 1e12));
//         for(int j = mid ; j >= 0 ; j--){
//             for(int k = 1 ; k >= 0 ; k--){
//                 ll ans = 1e12;
//                 ans = min(ans, pre[j][k]);

//                 if(k == 0 && j >= 1) ans = min(ans, tmp[i].first - tmp[i].second + pre[j-1][1]);
//                 else{
//                     if(j > 1) ans = min(ans, tmp[i].first + pre[j-1][1]);
//                     else ans = min(ans, tmp[i].first + tmp[i].second);
//                 }

//                 curr[j][k] = ans;
//             }
//         }
//         pre = curr;
//     }
//     // cout<<dp[0][mid][0]<<endl;
//     if(pre[mid][0] <= l) return 1;
//     return 0;
// }

// int main()
// {
//     fastio();
//     ll t;cin>>t;
//     while(t--){
//         ll n,l;cin>>n>>l;
//         vector<pair<ll,ll>> vp(n);
//         for (int i = 0; i < n; ++i)
//         {
//             cin>>vp[i].first>>vp[i].second;
//         }
//         vector<pair<ll,ll>> tmp = vp;
//         sort(vp.begin(), vp.end());
//         sort(tmp.begin(), tmp.end(), comp);

//         // for(int i = 0 ; i < n ; i++){
//         //     cout<<tmp[i].first<<' '<<tmp[i].second<<endl;
//         // }
//         // cout<<endl;

//         ll low = 1, high = n, ans = 0;
//         while(low <= high){
//             ll mid = (low + high) / 2;
//             if(check(tmp, mid, l)){
//                 ans = mid;
//                 low = mid+1;
//             }
//             else high = mid-1;
//         }
//         // check(tmp, 3, l);
//         cout<<ans<<endl;
//     }
// }

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

vector<vector<ll>> dp;

ll rec(ll ind, ll many, ll mid, vector<pair<ll, ll>> &tmp)
{
    if (ind == tmp.size())
        return 1e12;
    ll ans = 1e12;
    if (dp[ind][many] != -1)
        return dp[ind][many];
    // dont take
    // cout<<ind<<' '<<many<<endl;
    ans = min(ans, rec(ind + 1, many, mid, tmp));

    // take and move
    if (many == mid)
        ans = min(ans, tmp[ind].first - tmp[ind].second + rec(ind + 1, many - 1, mid, tmp));
    else if (many != mid)
    {
        if (many > 1)
            ans = min(ans, tmp[ind].first + rec(ind + 1, many - 1, mid, tmp));
        else
        {
            ans = min(ans, tmp[ind].first + tmp[ind].second);
        }
    }
    return dp[ind][many] = ans;
}

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
    dp.clear();
    dp.resize(n + 1, vector<ll>(mid + 1, -1));
    ll x = rec(0, mid, mid, tmp);
    // cout<<x<<endl;
    if (x <= l)
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