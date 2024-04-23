# Headstart: Effortless Project Structuring

This tool provides a variety of predefined project structures that you can choose from, allowing you to focus on what matters most - your code.

## Usage

### Installation

1. **Clone the Repository:**
   ```
   git clone https://github.com/mvstermind/headstart-go.git
   cd headstart-go
   ```

2. **Build and Install:**
   ```
   go build
   go list -f '{{.Target}}'
   export PATH=$PATH:/path/to/directory/you/copied
   go install
   ```

3. **Initialize Your Project:**
   ```
   headstartgo projectname
   ```

   Replace `projectname` with the desired name for the root of your project. Note that spaces are not allowed in the project name.

4. **Choose Your Template:**
   Select one of the provided templates based on your project's requirements.

5. **Have Fun Coding!**

## Folder Structures

### 1. Flat Structure
   - Suitable for small projects but may become harder to maintain in larger ones.

### 2. Layered Structure
   - Organizes components into layers for better separation of concerns.

### 3. Domain-Driven Design (DDD)
   - Ideal for large projects, organizes code around business domains.

### 4. Clean Architecture
   - Emphasizes a clear separation of concerns, promoting maintainability.

### 5. Modular Structure
   - Helpful for projects with multiple independent components.

Choose the structure that best fits your project's needs and start coding without worrying about organizing your files!

## Contributing

If you have suggestions for additional project structures or improvements to the tool, feel free to open an issue or submit a pull request. Your contributions are valuable!

## License

This project is licensed under the [MIT License](LICENSE), allowing you to modify and distribute the code freely, with proper attribution.

