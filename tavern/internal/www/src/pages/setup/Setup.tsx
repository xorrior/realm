import { useState, useEffect, FormEvent } from "react";

const API_ENDPOINT = process.env.REACT_APP_API_ENDPOINT ?? "";

export const Setup = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetch(`${API_ENDPOINT}/auth/setup`)
            .then((res) => res.json())
            .then((data) => {
                if (!data.needsSetup) {
                    window.location.href = "/login";
                }
                setLoading(false);
            })
            .catch(() => {
                setLoading(false);
            });
    }, []);

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        setError("");

        if (password !== confirmPassword) {
            setError("Passwords do not match");
            return;
        }

        if (password.length < 8) {
            setError("Password must be at least 8 characters");
            return;
        }

        if (username.length < 3 || username.length > 25) {
            setError("Username must be between 3 and 25 characters");
            return;
        }

        setLoading(true);

        try {
            const res = await fetch(`${API_ENDPOINT}/auth/setup`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ username, password }),
                credentials: "same-origin",
            });

            if (res.ok) {
                window.location.href = "/";
            } else {
                const data = await res.json().catch(() => ({}));
                setError(data.error || "Failed to create account");
            }
        } catch {
            setError("Failed to connect to server");
        } finally {
            setLoading(false);
        }
    };

    if (loading) {
        return (
            <div className="flex min-h-screen items-center justify-center bg-gray-900">
                <p className="text-gray-400">Loading...</p>
            </div>
        );
    }

    return (
        <div className="flex min-h-screen items-center justify-center bg-gray-900">
            <div className="w-full max-w-md rounded-lg bg-gray-800 p-8 shadow-lg">
                <h1 className="mb-2 text-center text-2xl font-bold text-white">
                    Realm Setup
                </h1>
                <p className="mb-6 text-center text-sm text-gray-400">
                    Create the first administrator account to get started.
                </p>
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label htmlFor="username" className="block text-sm font-medium text-gray-300">
                            Username
                        </label>
                        <input
                            id="username"
                            type="text"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            required
                            autoFocus
                            minLength={3}
                            maxLength={25}
                            className="mt-1 block w-full rounded-md border border-gray-600 bg-gray-700 px-3 py-2 text-white placeholder-gray-400 focus:border-purple-500 focus:outline-none focus:ring-1 focus:ring-purple-500"
                            placeholder="Choose a username"
                        />
                    </div>
                    <div>
                        <label htmlFor="password" className="block text-sm font-medium text-gray-300">
                            Password
                        </label>
                        <input
                            id="password"
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            required
                            minLength={8}
                            className="mt-1 block w-full rounded-md border border-gray-600 bg-gray-700 px-3 py-2 text-white placeholder-gray-400 focus:border-purple-500 focus:outline-none focus:ring-1 focus:ring-purple-500"
                            placeholder="Choose a password (min 8 characters)"
                        />
                    </div>
                    <div>
                        <label htmlFor="confirmPassword" className="block text-sm font-medium text-gray-300">
                            Confirm Password
                        </label>
                        <input
                            id="confirmPassword"
                            type="password"
                            value={confirmPassword}
                            onChange={(e) => setConfirmPassword(e.target.value)}
                            required
                            minLength={8}
                            className="mt-1 block w-full rounded-md border border-gray-600 bg-gray-700 px-3 py-2 text-white placeholder-gray-400 focus:border-purple-500 focus:outline-none focus:ring-1 focus:ring-purple-500"
                            placeholder="Confirm your password"
                        />
                    </div>
                    {error && (
                        <p className="text-sm text-red-400">{error}</p>
                    )}
                    <button
                        type="submit"
                        disabled={loading}
                        className="w-full rounded-md bg-purple-600 px-4 py-2 text-white hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2 focus:ring-offset-gray-800 disabled:opacity-50"
                    >
                        {loading ? "Creating account..." : "Create Account"}
                    </button>
                </form>
            </div>
        </div>
    );
};
