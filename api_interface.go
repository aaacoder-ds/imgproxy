package main

import "net/http"

var apiInterfaceTmpl = []byte(`
<!doctype html>
<html>
	<head>
		<title>API Documentation - imgproxy</title>
		<meta name="description" content="Complete API documentation for imgproxy - URL structure, processing options, and examples">
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
			.section {
				background: rgba(255,255,255,0.05);
				padding: 25px;
				border-radius: 8px;
				margin-bottom: 25px;
				border: 1px solid rgba(83, 209, 255, 0.2);
			}
			.section h2 {
				color: #53D1FF;
				margin-bottom: 20px;
				font-size: 1.8rem;
			}
			.section h3 {
				color: #53D1FF;
				margin: 20px 0 10px 0;
				font-size: 1.3rem;
			}
			.code-block {
				background: #1a1d26;
				padding: 20px;
				border-radius: 4px;
				margin: 15px 0;
				font-family: monospace;
				font-size: 14px;
				border: 1px solid #333;
				overflow-x: auto;
				line-height: 1.4;
			}
			.url-structure {
				background: linear-gradient(45deg, #53D1FF, #4bb8e6);
				color: #0d0f15;
				padding: 20px;
				border-radius: 8px;
				margin: 20px 0;
				font-family: monospace;
				font-size: 16px;
				font-weight: bold;
				text-align: center;
			}
			.option-table {
				width: 100%;
				border-collapse: collapse;
				margin: 15px 0;
			}
			.option-table th,
			.option-table td {
				padding: 12px;
				text-align: left;
				border-bottom: 1px solid #333;
			}
			.option-table th {
				background: rgba(83, 209, 255, 0.1);
				color: #53D1FF;
				font-weight: 600;
			}
			.option-table tr:hover {
				background: rgba(255,255,255,0.05);
			}
			.example-grid {
				display: grid;
				grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
				gap: 20px;
				margin: 20px 0;
			}
			.example-card {
				background: #1a1d26;
				padding: 20px;
				border-radius: 8px;
				border: 1px solid #333;
			}
			.example-card h4 {
				color: #53D1FF;
				margin-bottom: 10px;
			}
			.example-card p {
				color: #ccc;
				margin-bottom: 15px;
				font-size: 14px;
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
				margin: 5px;
				transition: background 0.3s;
			}
			.btn:hover {
				background: #4bb8e6;
			}
			.note {
				background: rgba(255, 193, 7, 0.1);
				border: 1px solid #ffc107;
				padding: 15px;
				border-radius: 4px;
				margin: 15px 0;
			}
			.note strong {
				color: #ffc107;
			}
			.warning {
				background: rgba(220, 53, 69, 0.1);
				border: 1px solid #dc3545;
				padding: 15px;
				border-radius: 4px;
				margin: 15px 0;
			}
			.warning strong {
				color: #dc3545;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>API Documentation</h1>
				<p>Complete reference for imgproxy API and processing options</p>
			</div>
			
			<div class="nav">
				<a href="/">Processing</a>
				<a href="/watermark">Watermark</a>
				<a href="/features">Features</a>
				<a href="/api" class="active">API Docs</a>
			</div>

			<div class="section">
				<h2>URL Structure</h2>
				<div class="url-structure">
					{signature}/{processing_options}/{source_url}
				</div>
				<p>The imgproxy URL consists of three main parts:</p>
				<ul style="margin: 15px 0; padding-left: 20px;">
					<li><strong>Signature:</strong> Security signature (use "unsafe" for testing)</li>
					<li><strong>Processing Options:</strong> Image transformation parameters</li>
					<li><strong>Source URL:</strong> Original image URL or encoded data</li>
				</ul>
			</div>

			<div class="section">
				<h2>Processing Options</h2>
				<p>Processing options are specified as URL parts divided by slashes. Each option has the format:</p>
				<div class="code-block">
					option_name:argument1:argument2:...:argumentN
				</div>

				<h3>Resize Options</h3>
				<table class="option-table">
					<thead>
						<tr>
							<th>Option</th>
							<th>Format</th>
							<th>Description</th>
						</tr>
					</thead>
					<tbody>
						<tr>
							<td>resize</td>
							<td>resize:type:width:height:enlarge:extend</td>
							<td>Resize image with specified parameters</td>
						</tr>
						<tr>
							<td>width</td>
							<td>width:value</td>
							<td>Set output width (0 = auto)</td>
						</tr>
						<tr>
							<td>height</td>
							<td>height:value</td>
							<td>Set output height (0 = auto)</td>
						</tr>
						<tr>
							<td>gravity</td>
							<td>gravity:position</td>
							<td>Set gravity for cropping/positioning</td>
						</tr>
					</tbody>
				</table>

				<h3>Watermark Options</h3>
				<table class="option-table">
					<thead>
						<tr>
							<th>Option</th>
							<th>Format</th>
							<th>Description</th>
						</tr>
					</thead>
					<tbody>
						<tr>
							<td>watermark_text</td>
							<td>watermark_text:text:font:color:opacity:gravity:offset_x:offset_y</td>
							<td>Add text watermark</td>
						</tr>
						<tr>
							<td>watermark_url</td>
							<td>watermark_url:url:size:opacity:gravity:offset_x:offset_y</td>
							<td>Add image watermark</td>
						</tr>
					</tbody>
				</table>

				<h3>Filter Options</h3>
				<table class="option-table">
					<thead>
						<tr>
							<th>Option</th>
							<th>Format</th>
							<th>Description</th>
						</tr>
					</thead>
					<tbody>
						<tr>
							<td>brightness</td>
							<td>brightness:value</td>
							<td>Adjust brightness (-100 to 100)</td>
						</tr>
						<tr>
							<td>contrast</td>
							<td>contrast:value</td>
							<td>Adjust contrast (-100 to 100)</td>
						</tr>
						<tr>
							<td>blur</td>
							<td>blur:value</td>
							<td>Apply blur effect (0 to 100)</td>
						</tr>
						<tr>
							<td>sharpen</td>
							<td>sharpen:value</td>
							<td>Apply sharpen effect (0 to 100)</td>
						</tr>
					</tbody>
				</table>

				<h3>Format Options</h3>
				<table class="option-table">
					<thead>
						<tr>
							<th>Option</th>
							<th>Format</th>
							<th>Description</th>
						</tr>
					</thead>
					<tbody>
						<tr>
							<td>format</td>
							<td>format:type</td>
							<td>Output format (jpg, png, webp, avif)</td>
						</tr>
						<tr>
							<td>quality</td>
							<td>quality:value</td>
							<td>Output quality (1 to 100)</td>
						</tr>
					</tbody>
				</table>
			</div>

			<div class="section">
				<h2>Source URL Formats</h2>
				
				<h3>Plain URL</h3>
				<div class="code-block">
					/plain/https://example.com/image.jpg
				</div>
				<p>Use for direct HTTP/HTTPS URLs. The URL should be percent-encoded.</p>

				<h3>Base64 Encoded</h3>
				<div class="code-block">
					/aHR0cDovL2V4YW1wbGUuY29tL2ltYWdlLmpwZw
				</div>
				<p>URL-safe Base64 encoding of the source URL.</p>

				<h3>Encrypted (Pro)</h3>
				<div class="code-block">
					/enc/encrypted_data_here
				</div>
				<p>AES-CBC encrypted source URL for enhanced security.</p>
			</div>

			<div class="section">
				<h2>Examples</h2>
				<div class="example-grid">
					<div class="example-card">
						<h4>Basic Resize</h4>
						<p>Resize image to 300x200 pixels</p>
						<div class="code-block">
							/unsafe/resize:fill:300:200:0/gravity:sm/plain/https://example.com/image.jpg
						</div>
						<a href="/" class="btn">Try This</a>
					</div>

					<div class="example-card">
						<h4>Add Watermark</h4>
						<p>Add text watermark to image</p>
						<div class="code-block">
							/unsafe/watermark_text:WATERMARK:48:ffffff:80:ce:0:0/plain/https://example.com/image.jpg
						</div>
						<a href="/watermark" class="btn">Try This</a>
					</div>

					<div class="example-card">
						<h4>Format Conversion</h4>
						<p>Convert to WebP with quality 85</p>
						<div class="code-block">
							/unsafe/format:webp/quality:85/plain/https://example.com/image.jpg
						</div>
						<a href="/" class="btn">Try This</a>
					</div>

					<div class="example-card">
						<h4>Apply Filters</h4>
						<p>Brightness +10, Contrast +20</p>
						<div class="code-block">
							/unsafe/brightness:10/contrast:20/plain/https://example.com/image.jpg
						</div>
						<a href="/" class="btn">Try This</a>
					</div>

					<div class="example-card">
						<h4>Complex Processing</h4>
						<p>Resize, watermark, and convert</p>
						<div class="code-block">
							/unsafe/resize:fill:800:600:0/watermark_text:PRO:36:ffffff:70:soea:20:20/format:webp/quality:90/plain/https://example.com/image.jpg
						</div>
						<a href="/" class="btn">Try This</a>
					</div>

					<div class="example-card">
						<h4>Base64 Source</h4>
						<p>Using Base64 encoded source URL</p>
						<div class="code-block">
							/unsafe/resize:fit:400:300:0/aHR0cDovL2V4YW1wbGUuY29tL2ltYWdlLmpwZw
						</div>
						<a href="/" class="btn">Try This</a>
					</div>
				</div>
			</div>

			<div class="section">
				<h2>Gravity Values</h2>
				<table class="option-table">
					<thead>
						<tr>
							<th>Value</th>
							<th>Position</th>
							<th>Description</th>
						</tr>
					</thead>
					<tbody>
						<tr><td>ce</td><td>Center</td><td>Center of the image</td></tr>
						<tr><td>no</td><td>North</td><td>Top center</td></tr>
						<tr><td>so</td><td>South</td><td>Bottom center</td></tr>
						<tr><td>ea</td><td>East</td><td>Right center</td></tr>
						<tr><td>we</td><td>West</td><td>Left center</td></tr>
						<tr><td>noea</td><td>North East</td><td>Top right corner</td></tr>
						<tr><td>nowe</td><td>North West</td><td>Top left corner</td></tr>
						<tr><td>soea</td><td>South East</td><td>Bottom right corner</td></tr>
						<tr><td>sowe</td><td>South West</td><td>Bottom left corner</td></tr>
						<tr><td>sm</td><td>Smart</td><td>Smart gravity (face detection)</td></tr>
					</tbody>
				</table>
			</div>

			<div class="section">
				<h2>Security Considerations</h2>
				
				<div class="warning">
					<strong>⚠️ Important:</strong> Always use URL signing in production environments. The "unsafe" signature is for testing only.
				</div>

				<h3>URL Signing</h3>
				<p>To sign URLs, you need to provide a key and salt pair to imgproxy. The signature is generated using HMAC-SHA256:</p>
				<div class="code-block">
					# Example in Go
					key := []byte("your-key")
					salt := []byte("your-salt")
					
					message := fmt.Sprintf("/%s/%s", processingOptions, sourceURL)
					h := hmac.New(sha256.New, key)
					h.Write(salt)
					h.Write([]byte(message))
					
					signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
				</div>

				<h3>Security Best Practices</h3>
				<ul style="margin: 15px 0; padding-left: 20px;">
					<li>Always use URL signing in production</li>
					<li>Set appropriate file size and resolution limits</li>
					<li>Whitelist allowed source domains</li>
					<li>Use HTTPS for all communications</li>
					<li>Regularly rotate keys and salts</li>
				</ul>
			</div>

			<div class="section">
				<h2>Error Handling</h2>
				<p>imgproxy returns appropriate HTTP status codes for different error conditions:</p>
				<table class="option-table">
					<thead>
						<tr>
							<th>Status Code</th>
							<th>Description</th>
						</tr>
					</thead>
					<tbody>
						<tr><td>200</td><td>Success</td></tr>
						<tr><td>400</td><td>Bad Request (invalid parameters)</td></tr>
						<tr><td>401</td><td>Unauthorized (invalid signature)</td></tr>
						<tr><td>404</td><td>Source image not found</td></tr>
						<tr><td>422</td><td>Processing failed</td></tr>
						<tr><td>500</td><td>Internal server error</td></tr>
					</tbody>
				</table>
			</div>

			<div class="section">
				<h2>Performance Tips</h2>
				<div class="note">
					<strong>💡 Tip:</strong> Use appropriate cache headers and CDN integration for optimal performance.
				</div>
				
				<h3>Optimization Strategies</h3>
				<ul style="margin: 15px 0; padding-left: 20px;">
					<li>Use WebP or AVIF for better compression</li>
					<li>Set appropriate quality values (85-90 for photos)</li>
					<li>Use progressive JPEG for large images</li>
					<li>Implement proper caching strategies</li>
					<li>Monitor memory usage and adjust limits</li>
				</ul>

				<h3>Recommended Settings</h3>
				<div class="code-block">
					# High quality photos
					format:webp/quality:90
					
					# Thumbnails
					format:webp/quality:80/resize:fill:150:150:0
					
					# Social media
					format:webp/quality:85/resize:fill:1200:630:0
				</div>
			</div>

			<div class="section">
				<h2>Get Started</h2>
				<p>Ready to start using imgproxy? Try our interactive interfaces:</p>
				<div style="text-align: center; margin: 20px 0;">
					<a href="/" class="btn">Processing Interface</a>
					<a href="/watermark" class="btn">Watermark Interface</a>
					<a href="/features" class="btn">Features Overview</a>
				</div>
			</div>
		</div>
	</body>
</html>
`)

func handleAPIInterface(reqID string, rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(200)
	rw.Write(apiInterfaceTmpl)
} 