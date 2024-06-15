ARG PARENT_CONTAINER_IMAGE_NAME="crypto-bundle/bc-wallet-common-trrfrm"

FROM $PARENT_CONTAINER_IMAGE_NAME

ARG TRFRM_SOURCE_DIR="/opt/trrfrm/source"
ARG TRFRM_PROJECT_NAME="bc-wallet-ethereum-hdwallet"
ENV TRFRM_PROJECT_NAME=$TRFRM_PROJECT_NAME

COPY deploy/trrfrm/hdwallet $TRFRM_SOURCE_DIR