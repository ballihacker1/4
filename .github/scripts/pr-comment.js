// const fs = require('fs');
// const path = require('path');

module.exports = async ({ github, context }) => {
  const { owner, repo } = context.repo;

  const prNumber = context.payload.pull_request.number;

  const sender = context.payload.sender.login;

  // Commant to be posted, Thanks to the sender
  const comment = `Thanks for the PR @${sender}!`;

  const { data: comments } = await github.issues.listComments({
    owner,
    repo,
    issue_number: prNumber,
  });

  const commentAlreadyExists = comments.find(
    (comment) => comment.body === comment,
  );

  if (commentAlreadyExists) {
    console.log(`The comment "${comment}" already exists`);
    return;
  }

  await github.rest.issues.createComment({
    owner,
    repo,
    issue_number: prNumber,
    body: comment,
  });
};
