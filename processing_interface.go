package main

import "net/http"

var processingInterfaceTmpl = []byte(`
<!doctype html>
<html>
	<head>
		<title>Image Processing Interface - imgproxy</title>
		<meta charset="utf-8">
		<meta name="description" content="Interactive image processing interface for imgproxy - resize, crop, filter, and transform images">
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
			.form-section {
				background: rgba(255,255,255,0.05);
				padding: 25px;
				border-radius: 8px;
				margin-bottom: 20px;
			}
			.form-section h3 {
				color: #53D1FF;
				margin-bottom: 15px;
				font-size: 1.3rem;
			}
			.form-group {
				margin-bottom: 15px;
			}
			.form-group label {
				display: block;
				margin-bottom: 5px;
				font-weight: 500;
			}
			.form-group input, .form-group select, .form-group textarea {
				width: 100%;
				padding: 10px;
				border: 1px solid #333;
				border-radius: 4px;
				background: #1a1d26;
				color: #fff;
				font-size: 14px;
			}
			.form-group input:focus, .form-group select:focus, .form-group textarea:focus {
				outline: none;
				border-color: #53D1FF;
			}
			.btn {
				background: #53D1FF;
				color: #0d0f15;
				padding: 12px 24px;
				border: none;
				border-radius: 4px;
				cursor: pointer;
				font-size: 16px;
				font-weight: 600;
				transition: background 0.3s;
			}
			.btn:hover {
				background: #4bb8e6;
			}
			.result-section {
				margin-top: 30px;
				text-align: center;
			}
			.result-image {
				max-width: 100%;
				border-radius: 8px;
				box-shadow: 0 4px 20px rgba(0,0,0,0.3);
			}
			.url-display {
				background: #1a1d26;
				padding: 15px;
				border-radius: 4px;
				margin: 20px 0;
				word-break: break-all;
				font-family: monospace;
				font-size: 12px;
				border: 1px solid #333;
			}
			.grid {
				display: grid;
				grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
				gap: 20px;
				margin-bottom: 20px;
			}
			.preset-btn {
				background: rgba(83, 209, 255, 0.1);
				border: 1px solid #53D1FF;
				color: #53D1FF;
				padding: 10px;
				border-radius: 4px;
				cursor: pointer;
				transition: all 0.3s;
			}
			.preset-btn:hover {
				background: rgba(83, 209, 255, 0.2);
			}
			.preset-btn.active {
				background: #53D1FF;
				color: #0d0f15;
			}
			.footer {
				position: fixed;
				bottom: 0;
				left: 0;
				width: 100%;
				background: rgba(13, 15, 21, 0.9);
				padding: 20px;
				text-align: center;
				font-size: 14px;
			}
			.footer a {
				color: #53D1FF;
				margin: 0 10px;
			}
			.main-content {
				margin-bottom: 100px;
			}
		</style>
	</head>
	<body>
		<div class="main-content">
			<div class="container">
			<div class="header">
				<h1>Image Processing Interface</h1>
				<p>Transform and optimize your images with imgproxy</p>
			</div>
			
			<div class="nav">
				<a href="/" class="active">Processing</a>
				<a href="/watermark">Watermark</a>
				<a href="/features">Features</a>
			</div>

			<form id="processingForm">
				<div class="form-section">
					<h3>Source Image</h3>
					<div class="form-group">
						<label for="sourceUrl">Image URL:</label>
						<input type="url" id="sourceUrl" name="sourceUrl" placeholder="https://example.com/image.jpg" required>
					</div>
				</div>

				<div class="form-section">
					<h3>Resize Options</h3>
					<div class="grid">
						<div class="form-group">
							<label for="resizeType">Resize Type:</label>
							<select id="resizeType" name="resizeType">
								<option value="fit">Fit (keep aspect ratio)</option>
								<option value="fill">Fill (crop to fit)</option>
								<option value="fill-down">Fill Down</option>
								<option value="force">Force (ignore aspect ratio)</option>
								<option value="auto">Auto</option>
							</select>
						</div>
						<div class="form-group">
							<label for="width">Width:</label>
							<input type="number" id="width" name="width" placeholder="300" min="0">
						</div>
						<div class="form-group">
							<label for="height">Height:</label>
							<input type="number" id="height" name="height" placeholder="400" min="0">
						</div>
						<div class="form-group">
							<label for="enlarge">Enlarge:</label>
							<select id="enlarge" name="enlarge">
								<option value="0">No</option>
								<option value="1">Yes</option>
							</select>
						</div>
					</div>
				</div>

				<div class="form-section">
					<h3>Quick Presets</h3>
					<div class="grid">
						<button type="button" class="preset-btn" data-preset="thumbnail">Thumbnail (150x150)</button>
						<button type="button" class="preset-btn" data-preset="medium">Medium (400x300)</button>
						<button type="button" class="preset-btn" data-preset="large">Large (800x600)</button>
						<button type="button" class="preset-btn" data-preset="square">Square (300x300)</button>
						<button type="button" class="preset-btn" data-preset="mobile">Mobile (320x480)</button>
						<button type="button" class="preset-btn" data-preset="desktop">Desktop (1920x1080)</button>
					</div>
				</div>

				<div class="form-section">
					<h3>Format & Quality</h3>
					<div class="grid">
						<div class="form-group">
							<label for="format">Output Format:</label>
							<select id="format" name="format">
								<option value="jpg">JPEG</option>
								<option value="png">PNG</option>
								<option value="webp">WebP</option>
								<option value="avif">AVIF</option>
							</select>
						</div>
						<div class="form-group">
							<label for="quality">Quality (1-100):</label>
							<input type="number" id="quality" name="quality" value="85" min="1" max="100">
						</div>
					</div>
				</div>

				<button type="submit" class="btn">Process Image</button>
			</form>

			<div class="result-section" id="resultSection" style="display: none;">
				<h3>Processed Image</h3>
				<div class="url-display" id="urlDisplay"></div>
				<img id="resultImage" class="result-image" alt="Processed image">
			</div>
		</div>
		</div>

		<div class="footer">
			<a href="/processing">Image Processing</a> |
			<a href="/watermark">Watermarking</a> |
			<a href="/features">Features</a> |
			<a href="https://imgproxy.net/" target="_blank">imgproxy.net</a> |
			<a href="https://dash.aaacoder.xyz/" target="_blank">Developer Tools</a>
		</div>

		<script defer data-domain="img.aaacoder.xyz" src="https://plausible.aaacoder.xyz/js/script.js"></script>
		<script id="aclib" type="text/javascript" src="//acscdn.com/script/aclib.js"></script>
		<script type="text/javascript">
			aclib.runAutoTag({
				zoneId: 'mhjbgr66iw',
			});
		</script>

		<script>
            document.getElementById('processingForm').addEventListener('submit', function(e) {
				e.preventDefault();
				processImage();
			});

			// Preset buttons
			document.querySelectorAll('.preset-btn').forEach(btn => {
				btn.addEventListener('click', function() {
					const preset = this.dataset.preset;
					applyPreset(preset);
				});
			});

			function applyPreset(preset) {
				const presets = {
					thumbnail: { width: 150, height: 150, resizeType: 'fill' },
					medium: { width: 400, height: 300, resizeType: 'fit' },
					large: { width: 800, height: 600, resizeType: 'fit' },
					square: { width: 300, height: 300, resizeType: 'fill' },
					mobile: { width: 320, height: 480, resizeType: 'fit' },
					desktop: { width: 1920, height: 1080, resizeType: 'fit' }
				};

				const config = presets[preset];
				if (config) {
					document.getElementById('width').value = config.width;
					document.getElementById('height').value = config.height;
					document.getElementById('resizeType').value = config.resizeType;
				}
			}

            function processImage() {
				const formData = new FormData(document.getElementById('processingForm'));
				const params = new URLSearchParams();
				
				// Build processing options
				const options = [];
				
				if (formData.get('resizeType') && formData.get('width') && formData.get('height')) {
					options.push('resize:' + formData.get('resizeType') + ':' + formData.get('width') + ':' + formData.get('height') + ':' + formData.get('enlarge'));
				}
				
				if (formData.get('quality')) {
					options.push('quality:' + formData.get('quality'));
				}
				
				if (formData.get('format')) {
					options.push('format:' + formData.get('format'));
				}

				// Build imgproxy URL
				const baseUrl = window.location.origin;
                const signature = 'unsafe'; // For demo purposes
				const processingOptions = options.join('/');
				const sourceUrl = encodeURIComponent(formData.get('sourceUrl'));
				const extension = formData.get('format') || 'jpg';
				
				const imgproxyUrl = baseUrl + '/' + signature + '/' + processingOptions + '/plain/' + sourceUrl + '@' + extension;
				
				// Display results
				document.getElementById('urlDisplay').textContent = imgproxyUrl;
				document.getElementById('resultImage').src = imgproxyUrl;
				document.getElementById('resultSection').style.display = 'block';
				
				// Scroll to results
				document.getElementById('resultSection').scrollIntoView({ behavior: 'smooth' });
                // Mark UI access for API calls
                document.cookie = "ui_access=1; path=/; max-age=86400";
			}
		</script>
	</body>
</html>
`)

func handleProcessingInterface(reqID string, rw http.ResponseWriter, r *http.Request) {
    rw.Header().Set("Content-Type", "text/html")
    rw.Header().Add("Set-Cookie", uiAccessCookieName+"=1; Path=/; Max-Age=86400")
	rw.WriteHeader(200)
	rw.Write(processingInterfaceTmpl)
} 