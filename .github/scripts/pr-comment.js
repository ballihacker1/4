// const fs = require('fs');
// const path = require('path');

module.exports = async ({ github, context }) => {
  const { owner, repo } = context.repo;

  const prNumber = context.payload.pull_request.number;

  const sender = context.payload.sender.login;

  const comment = `@${sender} Thanks for your PR!`;

  await github.rest.issues.createComment({
    owner,
    repo,
    issue_number: prNumber,
    body: comment,
  });
};
