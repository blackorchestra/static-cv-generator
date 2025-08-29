# CV Generator App

**Demo CV Template** - A static CV/resume generator that converts data from JSON to beautiful HTML with PDF export capability. This is a template project that you can fork and customize with your own information.

## 🚀 Features

- ✨ HTML generation from JSON data
- 📄 Export CV to PDF format
- 🎨 Responsive design
- 🔧 Easy to configure and modify
- 💼 Professional resume appearance
- 🚀 Automatic CI/CD via GitHub Actions
- ☁️ Hosting on AWS S3

## 📁 Project Structure

```
.
├── .github/
│   └── workflows/         # GitHub Actions for CI/CD
├── bin/                   # Compiled binary file
├── swp-generator/         # Go generator
│   ├── cmd/               
│   │   └── swpgenerator.go # Main program
│   ├── cv.json            # CV data
│   ├── template.html      # HTML template
│   └── go.mod             # Go module
└── web/                   # Web files
    ├── index.html         # Generated CV (do not edit manually!)
    ├── styles/            # CSS styles
    ├── scripts/           # JavaScript for PDF export
    └── images/            # Icons and images
```

## 🛠 Installation

### Requirements

- Go 1.18 or higher
- Git
- AWS CLI (for local deployment)

### Cloning the repository

```bash
git clone https://github.com/YOUR_USERNAME/cv-generator.git
cd cv-generator
```

## 📝 Usage

### 1. Editing CV Data

Edit the `swp-generator/cv.json` file with your personal information. The file contains demo data that you should replace:

```json
{
    "person_info": {
        "name": "Your name",
        "surname": "Your surname",
        "city": "Your city",
        "profile": "Brief description about yourself"
    },
    "position": "Desired position",
    "contact_information": {
        "phone_number": "Your phone",
        "email": "your.email@example.com",
        "github": "https://github.com/yourusername",
        "linkedin": "https://www.linkedin.com/in/yourprofile/"
    }
    // ... other fields
}
```

### 2. HTML Generation

#### Option 1: Using compiled binary

```bash
# Compilation (if needed)
go build -o bin/swpgenerator swp-generator/cmd/swpgenerator.go

# Run the generator
./bin/swpgenerator
```

#### Option 2: Run via Go

```bash
cd swp-generator
go run cmd/swpgenerator.go
```

#### Development Mode

For local development, use the `-environment=dev` flag:

```bash
cd swp-generator
go run cmd/swpgenerator.go -environment=dev
```

### 3. Viewing the Result

Open `web/index.html` in a browser to view the generated CV.

### 4. Export to PDF

On the CV page, click the "Export to PDF" button to download the CV in PDF format.

## 🚀 Deployment

You can deploy the generated HTML to various platforms:

### GitHub Pages
1. Generate the HTML with `./bin/swpgenerator`
2. Push to your repository
3. Enable GitHub Pages in repository settings

### AWS S3 (optional)
```bash
# Generate HTML
./bin/swpgenerator

# Sync with S3 (requires AWS CLI setup)
aws s3 sync web/ s3://your-bucket-name/ --delete
```

### Other Platforms
The `web/` directory contains static files that can be hosted on:
- Netlify
- Vercel
- Surge.sh
- Any static hosting service

## 🎨 Customization

### Changing CV Structure

1. Edit the data structure in `swp-generator/cmd/swpgenerator.go`
2. Update the template `swp-generator/template.html`
3. Adapt the data in `cv.json`

### Changing Styles

Edit CSS files in `web/styles/`:
- `normalize.css` - basic style normalization
- `style.css` - main CV styles

## 🔧 Development Commands

```bash
# Format Go code
cd swp-generator
go fmt ./...

# Check code
go vet ./...

# Build project
go build -o ../bin/swpgenerator cmd/swpgenerator.go
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Template

This is a demo template with sample data. Replace the information in `swp-generator/cv.json` with your own details to create your personal CV.

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve this CV generator.

---

*CV Template - Fork and customize with your own information* 📝
