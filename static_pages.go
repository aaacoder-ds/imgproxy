package main

import "net/http"

const uiAccessCookieName = "ui_access"

var notFoundTmpl = []byte(`
<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Page Not Found - imgproxy</title>
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<style>
			body { background:#0d0f15; color:#fff; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; }
			.container { max-width: 720px; margin: 10vh auto; padding: 0 20px; text-align: center; }
			h1 { font-size: 3rem; color:#53D1FF; margin-bottom: 10px; }
			p { color:#ccc; }
			.actions { margin-top: 30px; }
			.actions a { color:#0d0f15; background:#53D1FF; padding:10px 18px; border-radius:6px; text-decoration:none; margin: 0 8px; display:inline-block; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>404</h1>
			<p>We couldn't find that page.</p>
			<div class="actions">
				<a href="/">Home</a>
				<a href="/processing">Processing</a>
				<a href="/watermark">Watermark</a>
				<a href="/features">Features</a>
			</div>
		</div>
	</body>
</html>
`)

func handleNotFound(reqID string, rw http.ResponseWriter, r *http.Request) {
    rw.Header().Set("Content-Type", "text/html; charset=utf-8")
    rw.WriteHeader(http.StatusNotFound)
    rw.Write(notFoundTmpl)
}

var robotsTxt = []byte(`User-agent: *
Allow: /

# Allow access to all interface pages
Allow: /processing
Allow: /watermark
Allow: /features

Sitemap: https://img.aaacoder.xyz/sitemap.xml
`)

func handleRobotsTxt(reqID string, rw http.ResponseWriter, r *http.Request) {
    rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
    rw.WriteHeader(http.StatusOK)
    rw.Write(robotsTxt)
}

var sitemapXML = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://img.aaacoder.xyz/</loc>
    <lastmod>2024-01-01</lastmod>
    <changefreq>monthly</changefreq>
    <priority>1.0</priority>
  </url>
  <url>
    <loc>https://img.aaacoder.xyz/processing</loc>
    <lastmod>2024-01-01</lastmod>
    <changefreq>monthly</changefreq>
    <priority>0.9</priority>
  </url>
  <url>
    <loc>https://img.aaacoder.xyz/watermark</loc>
    <lastmod>2024-01-01</lastmod>
    <changefreq>monthly</changefreq>
    <priority>0.9</priority>
  </url>
  <url>
    <loc>https://img.aaacoder.xyz/features</loc>
    <lastmod>2024-01-01</lastmod>
    <changefreq>monthly</changefreq>
    <priority>0.8</priority>
  </url>
</urlset>`)

func handleSitemapXML(reqID string, rw http.ResponseWriter, r *http.Request) {
    rw.Header().Set("Content-Type", "application/xml; charset=utf-8")
    rw.WriteHeader(http.StatusOK)
    rw.Write(sitemapXML)
}


