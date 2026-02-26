import pandas as pd

def get_all_genres():
    """
    Loads movie data and extracts a sorted list of unique genres.
    """
    csv_path = 'dataes/movies_enriched.csv'
    try:
        df = pd.read_csv(csv_path)
    except FileNotFoundError:
        print(f"Error: The file {csv_path} was not found.")
        return []

    if 'genres' not in df.columns:
        print("Error: 'genres' column not in the dataframe.")
        return []

    # Handle missing values in 'genres' column
    df['genres'] = df['genres'].fillna('')

    # Split genres and find unique set
    all_genres = set()
    for s in df['genres']:
        if isinstance(s, str):
            all_genres.update(g for g in s.split('|') if g.strip() and g.strip() != '(no genres listed)')

    return sorted(list(all_genres))

if __name__ == '__main__':
    genres = get_all_genres()
    print(genres)
