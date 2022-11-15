import useSWR from 'swr';

const fetcher = (url: URL) => fetch(url).then(r => r.json())

export default function Home() {
  const { data, error } = useSWR('http://localhost:3001/lucky', fetcher)

  if (error) return <div>Hmm... noice error</div>
  if (!data) return <div>Hmm... loading...</div>
  return <div>{ data.luckyNumber }</div>
}
