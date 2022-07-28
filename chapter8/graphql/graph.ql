-------------------------------------------------------------------------------------------------------
query ($name: String!, $owner: String!) {
  repository(owner: $owner, name: $name) {
    createdAt
    forkCount
    labels(first: 5) {
      edges {
        node {
          name
        }
      }
    }
    issues(first: 5) {
      edges {
        node {
          title
        }
      }
    }
    commitComments(first: 10) {
      totalCount
      edges {
        node {
          author {
            url
            login
          }
        }
      }
    }
  }
}

-------------------------------------------------------------------------------------------------------

query ( $issueCommentsToAnalyze: Int!, $issuesToAnalyze: Int!, $name: String!, $owner: String!) {
  repository(owner: $owner, name: $name) {
    issues(first: $issuesToAnalyze, orderBy: {field: UPDATED_AT, direction: DESC}) {
      nodes {
        url
        authorAssociation
        author {
          login
        }
        createdAt
        comments(last: $issueCommentsToAnalyze) {
          nodes {
            authorAssociation
            createdAt
            author {
              login
            }
          }
        }
      }
    }
  }
  rateLimit {
    cost
  }
}
