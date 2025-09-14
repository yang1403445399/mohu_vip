<script setup lang="ts"></script>

<template>
  <router-view />
</template>

<style lang="scss">
* {
  font-family: "Inter Var", Inter, -apple-system, BlinkMacSystemFont,
    "San Francisco", "Segoe UI", Roboto, "Helvetica Neue", sans-serif;
}

.el-statistic {
  .el-statistic__content {
    line-height: 1;
    display: flex;
    align-items: flex-end;

    .el-statistic__number {
      font-weight: bold;
      font-size: 30px;
    }
  }
}

.el-table {
  border-top: 1px solid var(--el-border-color);

  thead {
    tr {
      th {
        &.el-table__cell {
          background: var(--grey-bg-color);

          .cell {
            padding: 0 16px;
            display: flex;
            align-items: center;
          }
        }
      }
    }
  }

  tbody {
    tr {
      td {
        &.el-table__cell {
          .cell {
            padding: 0 16px;
            display: flex;
            align-items: center;
          }
        }
      }
    }
  }
}

.el-dialog {
  padding: 0 !important;
  overflow: hidden;

  .el-dialog__header {
    padding: 16px !important;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid var(--el-border-color);

    .el-dialog__title {
      line-height: 1;
      font-size: 16px;
    }

    .el-dialog__headerbtn {
      position: static;
      height: 16px;
      width: 16px;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }

  .el-dialog__footer {
    padding: 0 16px 16px 16px;
  }
}

.el-message-box {
  padding: 0 !important;

  .el-message-box__header {
    padding: 16px !important;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .el-message-box__title {
      font-size: 16px;
      line-height: 1;
    }

    .el-message-box__headerbtn {
      position: static;
      height: 16px;
      width: 16px;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }

  .el-message-box__container {
    padding: 8px 16px 0 16px;

    .el-message-box__status {
      display: flex;
    }

    .el-message-box__message {
      p {
        line-height: 1;
      }
    }
  }

  .el-message-box__btns {
    padding: 16px;
  }
}

.el-pagination {
  .el-pagination__total {
    line-height: 1;
  }

  .el-pagination__sizes {
    display: flex;
  }

  .el-pagination__jump {
    line-height: 1;
  }
}

.el-input-number {
  &.is-controls-right {
    .el-input__wrapper {
      padding: 1px 11px !important;
      .el-input__inner {
        text-align: left;
        line-height: 30px;
      }
    }
  }
}
</style>
