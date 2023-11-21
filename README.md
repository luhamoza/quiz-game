# Quiz Game

This command-line quiz game reads questions and answers from a CSV file, allowing users to test their knowledge. The default CSV file is "problems.csv," but users can specify another file using the `-csv` flag.

## Usage

```bash
go run main.go -csv your_questions.csv
```

## CSV Format

The CSV file should follow the format: `question,answer`.

## How to Play

1. Run the program.
2. Answer questions by typing your response and pressing Enter.
3. Get your score at the end.

## Example

```bash
go run main.go -csv sample_questions.csv
```

Ensure your CSV file is properly formatted to avoid parsing errors.
