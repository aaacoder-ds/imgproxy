package main

import "net/http"

var featuresInterfaceTmpl = []byte(`
<!doctype html>
<html>
	<head>
		<title>Features Overview - imgproxy</title>
		<meta name="description" content="Comprehensive overview of imgproxy features - image processing, optimization, transformation, and more">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<style>
			* {
				margin: 0;
				padding: 0;
				box-sizing: border-box;
			}
			body {
				font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
				background: #0d0f15;
				color: #fff;
				line-height: 1.6;
			}
			.container {
				max-width: 1200px;
				margin: 0 auto;
				padding: 20px;
			}
			.header {
				text-align: center;
				margin-bottom: 40px;
			}
			.header h1 {
				font-size: 2.5rem;
				margin-bottom: 10px;
				color: #53D1FF;
			}
			.header p {
				font-size: 1.2rem;
				color: #ccc;
			}
			.nav {
				background: rgba(255,255,255,0.1);
				padding: 15px;
				border-radius: 8px;
				margin-bottom: 30px;
			}
			.nav a {
				color: #53D1FF;
				text-decoration: none;
				margin-right: 20px;
				padding: 8px 16px;
				border-radius: 4px;
				transition: background 0.3s;
			}
			.nav a:hover {
				background: rgba(83, 209, 255, 0.2);
			}
			.nav a.active {
				background: #53D1FF;
				color: #0d0f15;
			}
			.feature-grid {
				display: grid;
				grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
				gap: 25px;
				margin-bottom: 40px;
			}
			.feature-card {
				background: rgba(255,255,255,0.05);
				padding: 25px;
				border-radius: 8px;
				border: 1px solid rgba(83, 209, 255, 0.2);
				transition: all 0.3s;
			}
			.feature-card:hover {
				border-color: #53D1FF;
				transform: translateY(-2px);
			}
			.feature-card h3 {
				color: #53D1FF;
				margin-bottom: 15px;
				font-size: 1.4rem;
			}
			.feature-card p {
				color: #ccc;
				margin-bottom: 15px;
			}
			.feature-list {
				list-style: none;
				margin-top: 15px;
			}
			.feature-list li {
				padding: 5px 0;
				color: #ddd;
				position: relative;
				padding-left: 20px;
			}
			.feature-list li:before {
				content: "✓";
				color: #53D1FF;
				position: absolute;
				left: 0;
				font-weight: bold;
			}
			.code-example {
				background: #1a1d26;
				padding: 15px;
				border-radius: 4px;
				margin: 15px 0;
				font-family: monospace;
				font-size: 12px;
				border: 1px solid #333;
				overflow-x: auto;
			}
			.btn {
				background: #53D1FF;
				color: #0d0f15;
				padding: 10px 20px;
				border: none;
				border-radius: 4px;
				cursor: pointer;
				font-size: 14px;
				font-weight: 600;
				text-decoration: none;
				display: inline-block;
				margin-top: 10px;
				transition: background 0.3s;
			}
			.btn:hover {
				background: #4bb8e6;
			}
			.stats-section {
				background: rgba(83, 209, 255, 0.1);
				padding: 30px;
				border-radius: 8px;
				margin: 40px 0;
				text-align: center;
			}
			.stats-grid {
				display: grid;
				grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
				gap: 20px;
				margin-top: 20px;
			}
			.stat-item {
				text-align: center;
			}
			.stat-number {
				font-size: 2.5rem;
				font-weight: bold;
				color: #53D1FF;
			}
			.stat-label {
				color: #ccc;
				font-size: 1rem;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>Features Overview</h1>
				<p>Discover the powerful capabilities of imgproxy</p>
			</div>
			
			<div class="nav">
				<a href="/">Processing</a>
				<a href="/watermark">Watermark</a>
				<a href="/features" class="active">Features</a>
				<a href="/api">API Docs</a>
			</div>

			<div class="stats-section">
				<h2>Why Choose imgproxy?</h2>
				<div class="stats-grid">
					<div class="stat-item">
						<div class="stat-number">50+</div>
						<div class="stat-label">Processing Options</div>
					</div>
					<div class="stat-item">
						<div class="stat-number">10+</div>
						<div class="stat-label">Image Formats</div>
					</div>
					<div class="stat-item">
						<div class="stat-number">1000+</div>
						<div class="stat-label">Images/Second</div>
					</div>
					<div class="stat-item">
						<div class="stat-number">99.9%</div>
						<div class="stat-label">Uptime</div>
					</div>
				</div>
			</div>

			<div class="feature-grid">
				<div class="feature-card">
					<h3>🖼️ Image Resizing</h3>
					<p>Resize images with multiple algorithms and maintain aspect ratios</p>
					<ul class="feature-list">
						<li>Fit, Fill, Force, and Auto resize modes</li>
						<li>Multiple resizing algorithms (Lanczos, Cubic, Linear)</li>
						<li>Smart gravity positioning</li>
						<li>Enlarge and extend options</li>
					</ul>
					<div class="code-example">
						resize:fill:800:600:0/gravity:sm
					</div>
					<a href="/" class="btn">Try Resizing</a>
				</div>

				<div class="feature-card">
					<h3>💧 Watermarking</h3>
					<p>Add text and image watermarks with precise positioning</p>
					<ul class="feature-list">
						<li>Text watermarks with custom fonts</li>
						<li>Image watermarks with opacity control</li>
						<li>9-position gravity system</li>
						<li>Custom offset positioning</li>
					</ul>
					<div class="code-example">
						watermark_text:WATERMARK:48:ffffff:80:ce:0:0
					</div>
					<a href="/watermark" class="btn">Try Watermarking</a>
				</div>

				<div class="feature-card">
					<h3>🎨 Image Filters</h3>
					<p>Apply various filters and effects to enhance images</p>
					<ul class="feature-list">
						<li>Brightness, Contrast, Saturation</li>
						<li>Blur and Sharpen effects</li>
						<li>Pixelate and Unsharp masking</li>
						<li>Colorize and Gradient overlays</li>
					</ul>
					<div class="code-example">
						brightness:10/contrast:20/blur:5
					</div>
					<a href="/" class="btn">Try Filters</a>
				</div>

				<div class="feature-card">
					<h3>✂️ Cropping & Trimming</h3>
					<p>Precise image cropping and automatic trimming</p>
					<ul class="feature-list">
						<li>Smart cropping with gravity</li>
						<li>Automatic background trimming</li>
						<li>Custom crop dimensions</li>
						<li>Padding and extension options</li>
					</ul>
					<div class="code-example">
						crop:300:200:ce:0:0/trim:10
					</div>
					<a href="/" class="btn">Try Cropping</a>
				</div>

				<div class="feature-card">
					<h3>🔄 Format Conversion</h3>
					<p>Convert between multiple image formats with optimization</p>
					<ul class="feature-list">
						<li>JPEG, PNG, WebP, AVIF support</li>
						<li>Automatic format selection</li>
						<li>Quality and compression control</li>
						<li>Progressive JPEG support</li>
					</ul>
					<div class="code-example">
						format:webp/quality:85
					</div>
					<a href="/" class="btn">Try Conversion</a>
				</div>

				<div class="feature-card">
					<h3>🛡️ Security Features</h3>
					<p>Built-in security to protect your image processing</p>
					<ul class="feature-list">
						<li>URL signing and encryption</li>
						<li>Source URL validation</li>
						<li>File size and resolution limits</li>
						<li>Allowed domains whitelist</li>
					</ul>
					<div class="code-example">
						max_src_resolution:268402689/max_src_file_size:10485760
					</div>
					<a href="/api" class="btn">Learn Security</a>
				</div>

				<div class="feature-card">
					<h3>⚡ Performance</h3>
					<p>Optimized for high-performance image processing</p>
					<ul class="feature-list">
						<li>LibVIPS backend for speed</li>
						<li>Memory-efficient processing</li>
						<li>Concurrent request handling</li>
						<li>CDN-friendly caching</li>
					</ul>
					<div class="code-example">
						# Processes 1000+ images per second
					</div>
					<a href="/api" class="btn">View Performance</a>
				</div>

				<div class="feature-card">
					<h3>🌐 Multiple Sources</h3>
					<p>Support for various image sources and storage</p>
					<ul class="feature-list">
						<li>HTTP/HTTPS URLs</li>
						<li>Amazon S3 integration</li>
						<li>Google Cloud Storage</li>
						<li>Azure Blob Storage</li>
					</ul>
					<div class="code-example">
						# Works with any accessible image URL
					</div>
					<a href="/api" class="btn">Learn Sources</a>
				</div>

				<div class="feature-card">
					<h3>📊 Monitoring & Metrics</h3>
					<p>Comprehensive monitoring and analytics</p>
					<ul class="feature-list">
						<li>Prometheus metrics</li>
						<li>Health check endpoints</li>
						<li>Error reporting integration</li>
						<li>Performance monitoring</li>
					</ul>
					<div class="code-example">
						# Built-in /health endpoint
					</div>
					<a href="/api" class="btn">View Metrics</a>
				</div>
			</div>

			<div class="feature-card" style="grid-column: 1 / -1;">
				<h3>🚀 Getting Started</h3>
				<p>Ready to start using imgproxy? Here's a simple example:</p>
				<div class="code-example">
					# Basic resize example
					https://imgproxy.example.com/unsafe/resize:fill:300:200:0/gravity:sm/plain/https://example.com/image.jpg@webp
				</div>
				<p style="margin-top: 15px;">
					<a href="/" class="btn">Try Processing Interface</a>
					<a href="/api" class="btn">View API Documentation</a>
				</p>
			</div>
		</div>
	</body>
</html>
`)

func handleFeaturesInterface(reqID string, rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(200)
	rw.Write(featuresInterfaceTmpl)
} 