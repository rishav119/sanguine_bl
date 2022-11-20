import { Link } from 'react-router-dom'

export function PageLink({ text, url, external = false }) {
  if (external) {
    return (
      <div className="mt-2 mb-10 ml-2">
        <a
          className="text-white text-opacity-50 hover:text-opacity-90 hover:underline"
          href={url}
          target="_blank"
          rel="noreferrer"
        >
          {text}
        </a>
      </div>
    )
  } else {
    return (
      <div className="mt-2 mb-10 ml-2">
        <Link
          className="text-white text-opacity-50 hover:text-opacity-90 hover:underline"
          to={url}
        >
          {text}
        </Link>
      </div>
    )
  }
}
