"use client";

import { useState } from "react";

type EloEntry = {
  p1_elo: number;
  tourney_id: string;
  match_num: number;
};

export default function Home() {
  const [playerId, setPlayerId] = useState("");
  const [playerName, setPlayerName] = useState("");
  const [entries, setEntries] = useState<EloEntry[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [searchedId, setSearchedId] = useState<string | null>(null);

  console.log("test")

  const fetchElos = async (e: React.FormEvent) => {

    e.preventDefault();
    if (!playerId.trim()) return;

    setLoading(true);
    setError(null);
    setEntries([]);
    setSearchedId(null);

    try {
      const res = await fetch(`/api/elos/${playerId.trim()}`);
      const res2 = await fetch(`/api/player/${playerId.trim()}`);
      const data2 = await res2.json();
      if (!res.ok) {
        throw new Error("Failed to fetch player data.");
      }
      const data = await res.json();
      if (data.entries && data.entries.length > 0) {
        setEntries(data.entries);
        setSearchedId(playerId.trim());
        setPlayerName(data2.player_name)
      } else {
        setError("No Elo records found for this player ID.");
      }
    } catch (err: any) {
      setError(err.message || "An unexpected error occurred.");
    } finally {
      setLoading(false);
    }
  };

  const fetchPlayerName = async (e: React.FormEvent) => {
      const res = await fetch(`/api/player/${playerId.trim()}}`);

      const data = await res.json()
  }

  return (
    <div className="min-h-screen bg-zinc-950 text-zinc-50 font-sans selection:bg-indigo-500/30">
      <div className="absolute inset-0 bg-[radial-gradient(ellipse_at_top,_var(--tw-gradient-stops))] from-indigo-900/20 via-zinc-950 to-zinc-950 -z-10" />

      <main className="container mx-auto max-w-4xl px-6 py-24">
        <div className="flex flex-col items-center text-center space-y-8 mb-16">
          <div className="inline-flex items-center justify-center p-2 bg-indigo-500/10 rounded-full mb-4 border border-indigo-500/20">
            <span className="text-xs font-semibold tracking-wider text-indigo-400 uppercase px-3 py-1">
              Tennis Analytics
            </span>
          </div>

          <h1 className="text-5xl md:text-7xl font-extrabold tracking-tight text-transparent bg-clip-text bg-gradient-to-br from-white via-white to-zinc-500 pb-2">
            Player Elo Ratings
          </h1>

          <p className="max-w-2xl text-lg md:text-xl text-zinc-400">
            Enter an ATP player ID to visualize their historical match Elo ratings across different tournaments.
          </p>

          <form onSubmit={fetchElos} className="w-full max-w-md relative group mt-8">
            <div className="absolute -inset-0.5 bg-gradient-to-r from-indigo-500 to-purple-500 rounded-2xl blur opacity-30 group-hover:opacity-60 transition duration-500"></div>
            <div className="relative flex items-center bg-zinc-900 rounded-2xl p-2 ring-1 ring-white/10 shadow-2xl">
              <input
                type="text"
                placeholder="Enter Player ID (e.g. 206173)"
                value={playerId}
                onChange={(e) => setPlayerId(e.target.value)}
                className="w-full bg-transparent border-none px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:ring-0 text-lg"
                required
              />
              <div>
                {playerName}
              </div>
              <button
                type="submit"
                disabled={loading}
                className="ml-2 bg-indigo-600 hover:bg-indigo-500 text-white font-medium px-6 py-3 rounded-xl transition-all active:scale-95 disabled:opacity-50 flex items-center justify-center min-w-[120px]"
              >
                {loading ? (
                  <svg className="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                    <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                    <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                ) : (
                  "Search"
                )}
              </button>
            </div>
          </form>
        </div>

        {error && (
          <div className="max-w-md mx-auto mt-8 p-4 bg-red-500/10 border border-red-500/20 rounded-2xl text-red-400 text-center animate-in fade-in slide-in-from-bottom-4 duration-500">
            {error}
          </div>
        )}

        {entries.length > 0 && !loading && (
          <div className="animate-in fade-in slide-in-from-bottom-8 duration-700 mt-16">
            <div className="flex items-center justify-between mb-8 px-2">
              <h2 className="text-2xl font-semibold">Results for Player #{searchedId}</h2>
              <span className="text-zinc-400 bg-zinc-800/50 px-4 py-1.5 rounded-full text-sm font-medium border border-zinc-700/50">
                {entries.length} matches found
              </span>
            </div>

            <div className="bg-zinc-900/50 backdrop-blur-xl border border-zinc-800 rounded-3xl overflow-hidden shadow-2xl">
              <div className="overflow-x-auto">
                <table className="w-full text-left border-collapse">
                  <thead>
                    <tr className="bg-zinc-800/50 text-zinc-300 text-sm uppercase tracking-wider">
                      <th className="px-8 py-5 font-semibold">Tournament ID</th>
                      <th className="px-8 py-5 font-semibold">Match Num</th>
                      <th className="px-8 py-5 font-semibold text-right">Player Elo</th>
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-zinc-800/50">
                    {entries.map((entry, index) => (
                      <tr
                        key={index}
                        className="hover:bg-zinc-800/30 transition-colors group"
                      >
                        <td className="px-8 py-5 text-zinc-300 font-medium group-hover:text-white transition-colors">
                          {entry.tourney_id}
                        </td>
                        <td className="px-8 py-5 text-zinc-400">
                          {entry.match_num}
                        </td>
                        <td className="px-8 py-5 text-right font-mono font-medium text-indigo-300 group-hover:text-indigo-200 transition-colors">
                          {entry.p1_elo.toFixed(2)}
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        )}
      </main>
    </div>
  );
}
