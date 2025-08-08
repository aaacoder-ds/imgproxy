package main

import "net/http"

var watermarkInterfaceTmpl = []byte(`
<!doctype html>
<html>
	<head>
		<title>Watermark Interface - imgproxy</title>
		<meta charset="utf-8">
		<meta name="description" content="Add watermarks to images with imgproxy - text and image watermarks with positioning and styling options">
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
			.watermark-type {
				background: rgba(83, 209, 255, 0.1);
				border: 1px solid #53D1FF;
				color: #53D1FF;
				padding: 10px;
				border-radius: 4px;
				cursor: pointer;
				transition: all 0.3s;
				text-align: center;
			}
			.watermark-type:hover {
				background: rgba(83, 209, 255, 0.2);
			}
			.watermark-type.active {
				background: #53D1FF;
				color: #0d0f15;
			}
			.preview-box {
				border: 2px dashed #53D1FF;
				padding: 20px;
				border-radius: 8px;
				margin: 20px 0;
				text-align: center;
				min-height: 100px;
				display: flex;
				align-items: center;
				justify-content: center;
			}
			.color-picker {
				width: 50px;
				height: 40px;
				border: none;
				border-radius: 4px;
				cursor: pointer;
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
				<h1>Watermark Interface</h1>
				<p>Add text and image watermarks to your images with imgproxy</p>
			</div>
			
			<div class="nav">
				<a href="/">Processing</a>
				<a href="/watermark" class="active">Watermark</a>
				<a href="/features">Features</a>
			</div>

			<form id="watermarkForm">
				<div class="form-section">
					<h3>Source Image</h3>
					<div class="form-group">
						<label for="sourceUrl">Image URL:</label>
						<input type="url" id="sourceUrl" name="sourceUrl" placeholder="https://example.com/image.jpg" required>
					</div>
				</div>

				<div class="form-section">
					<h3>Watermark Type</h3>
					<div class="grid">
						<div class="watermark-type active" data-type="text">
							<h4>Text Watermark</h4>
							<p>Add text overlay</p>
						</div>
						<div class="watermark-type" data-type="image">
							<h4>Image Watermark</h4>
							<p>Add image overlay</p>
						</div>
					</div>
				</div>

				<div class="form-section" id="textWatermarkSection">
					<h3>Text Watermark Settings</h3>
					<div class="grid">
						<div class="form-group">
							<label for="watermarkText">Watermark Text:</label>
							<input type="text" id="watermarkText" name="watermarkText" placeholder="Your watermark text" value="WATERMARK">
						</div>
						<div class="form-group">
							<label for="watermarkFont">Font Size:</label>
							<input type="number" id="watermarkFont" name="watermarkFont" value="48" min="8" max="200">
						</div>
						<div class="form-group">
							<label for="watermarkColor">Text Color:</label>
							<input type="color" id="watermarkColor" name="watermarkColor" value="#ffffff" class="color-picker">
						</div>
						<div class="form-group">
							<label for="watermarkOpacity">Opacity (0-100):</label>
							<input type="number" id="watermarkOpacity" name="watermarkOpacity" value="80" min="0" max="100">
						</div>
					</div>
				</div>

				<div class="form-section" id="imageWatermarkSection" style="display: none;">
					<h3>Image Watermark Settings</h3>
					<div class="grid">
						<div class="form-group">
							<label for="watermarkImageUrl">Watermark Image URL:</label>
							<input type="url" id="watermarkImageUrl" name="watermarkImageUrl" placeholder="https://example.com/watermark.png">
						</div>
						<div class="form-group">
							<label for="watermarkSize">Size (% of image):</label>
							<input type="number" id="watermarkSize" name="watermarkSize" value="20" min="1" max="100">
						</div>
						<div class="form-group">
							<label for="watermarkOpacity">Opacity (0-100):</label>
							<input type="number" id="watermarkOpacity" name="watermarkOpacity" value="80" min="0" max="100">
						</div>
					</div>
				</div>

				<div class="form-section">
					<h3>Positioning</h3>
					<div class="grid">
						<div class="form-group">
							<label for="watermarkGravity">Position:</label>
							<select id="watermarkGravity" name="watermarkGravity">
								<option value="ce">Center</option>
								<option value="no">North</option>
								<option value="so">South</option>
								<option value="ea">East</option>
								<option value="we">West</option>
								<option value="noea">North East</option>
								<option value="nowe">North West</option>
								<option value="soea">South East</option>
								<option value="sowe">South West</option>
							</select>
						</div>
						<div class="form-group">
							<label for="watermarkOffsetX">X Offset:</label>
							<input type="number" id="watermarkOffsetX" name="watermarkOffsetX" value="0">
						</div>
						<div class="form-group">
							<label for="watermarkOffsetY">Y Offset:</label>
							<input type="number" id="watermarkOffsetY" name="watermarkOffsetY" value="0">
						</div>
					</div>
				</div>

				<div class="form-section">
					<h3>Preview</h3>
					<div class="preview-box" id="previewBox">
						<p>Watermark preview will appear here</p>
					</div>
				</div>

				<button type="submit" class="btn">Add Watermark</button>
			</form>

			<div class="result-section" id="resultSection" style="display: none;">
				<h3>Watermarked Image</h3>
				<div class="url-display" id="urlDisplay"></div>
				<img id="resultImage" class="result-image" alt="Watermarked image">
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
			let currentWatermarkType = 'text';

			// Watermark type selection
			document.querySelectorAll('.watermark-type').forEach(btn => {
				btn.addEventListener('click', function() {
					document.querySelectorAll('.watermark-type').forEach(b => b.classList.remove('active'));
					this.classList.add('active');
					currentWatermarkType = this.dataset.type;
					
					if (currentWatermarkType === 'text') {
						document.getElementById('textWatermarkSection').style.display = 'block';
						document.getElementById('imageWatermarkSection').style.display = 'none';
					} else {
						document.getElementById('textWatermarkSection').style.display = 'none';
						document.getElementById('imageWatermarkSection').style.display = 'block';
					}
					updatePreview();
				});
			});

			// Update preview on input changes
			document.querySelectorAll('#watermarkForm input, #watermarkForm select').forEach(input => {
				input.addEventListener('input', updatePreview);
			});

			function updatePreview() {
				const previewBox = document.getElementById('previewBox');
				const formData = new FormData(document.getElementById('watermarkForm'));
				
				if (currentWatermarkType === 'text') {
					const text = formData.get('watermarkText') || 'WATERMARK';
					const color = formData.get('watermarkColor') || '#ffffff';
					const opacity = formData.get('watermarkOpacity') || 80;
					
					previewBox.innerHTML = '<div style="color: ' + color + '; opacity: ' + (opacity/100) + '; font-size: 24px; font-weight: bold;">' + text + '</div>';
				} else {
					const imageUrl = formData.get('watermarkImageUrl');
					if (imageUrl) {
						previewBox.innerHTML = '<img src="' + imageUrl + '" style="max-width: 100px; max-height: 100px; opacity: ' + (formData.get('watermarkOpacity')/100 || 0.8) + ';">';
					} else {
						previewBox.innerHTML = '<p>Enter watermark image URL to see preview</p>';
					}
				}
			}

			document.getElementById('watermarkForm').addEventListener('submit', function(e) {
				e.preventDefault();
				addWatermark();
			});

            function addWatermark() {
				const formData = new FormData(document.getElementById('watermarkForm'));
				const options = [];
				
				// Basic resize for demo
				options.push('resize:fit:800:600:0');
				
				if (currentWatermarkType === 'text') {
					const text = formData.get('watermarkText');
					const font = formData.get('watermarkFont');
					const color = formData.get('watermarkColor').replace('#', '');
					const opacity = formData.get('watermarkOpacity');
					const gravity = formData.get('watermarkGravity');
					const offsetX = formData.get('watermarkOffsetX');
					const offsetY = formData.get('watermarkOffsetY');
					
					if (text) {
						options.push('watermark_text:' + text + ':' + font + ':' + color + ':' + opacity + ':' + gravity + ':' + offsetX + ':' + offsetY);
					}
				} else {
					const imageUrl = formData.get('watermarkImageUrl');
					const size = formData.get('watermarkSize');
					const opacity = formData.get('watermarkOpacity');
					const gravity = formData.get('watermarkGravity');
					const offsetX = formData.get('watermarkOffsetX');
					const offsetY = formData.get('watermarkOffsetY');
					
					if (imageUrl) {
						options.push('watermark_url:' + encodeURIComponent(imageUrl) + ':' + size + ':' + opacity + ':' + gravity + ':' + offsetX + ':' + offsetY);
					}
				}

				// Build imgproxy URL
				const baseUrl = window.location.origin;
                const signature = 'unsafe';
				const processingOptions = options.join('/');
				const sourceUrl = encodeURIComponent(formData.get('sourceUrl'));
				
				const imgproxyUrl = baseUrl + '/' + signature + '/' + processingOptions + '/plain/' + sourceUrl + '@jpg';
				
				// Display results
				document.getElementById('urlDisplay').textContent = imgproxyUrl;
				document.getElementById('resultImage').src = imgproxyUrl;
				document.getElementById('resultSection').style.display = 'block';
				
				// Scroll to results
				document.getElementById('resultSection').scrollIntoView({ behavior: 'smooth' });
                // Mark UI access for API calls
                document.cookie = "ui_access=1; path=/; max-age=86400";
			}

			// Initialize preview
            updatePreview();
		</script>
	</body>
</html>
`)

func handleWatermarkInterface(reqID string, rw http.ResponseWriter, r *http.Request) {
    rw.Header().Set("Content-Type", "text/html")
    rw.Header().Add("Set-Cookie", uiAccessCookieName+"=1; Path=/; Max-Age=86400")
	rw.WriteHeader(200)
	rw.Write(watermarkInterfaceTmpl)
} 